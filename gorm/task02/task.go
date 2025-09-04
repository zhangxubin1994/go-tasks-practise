package task02

import (
	"fmt"
	"gorm.io/gorm"
	"log"
)

/**
  sql语句练习 - 事务语句
*/

// 账户模型
type Account struct {
	ID      uint `gorm:"primaryKey"`
	Balance float64
}

// 交易记录模型
type Transaction struct {
	ID            uint `gorm:"primaryKey"`
	FromAccountID uint
	ToAccountID   uint
	Amount        float64
}

func Add(db *gorm.DB) {
	// 自动迁移表结构
	err := db.AutoMigrate(&Account{}, &Transaction{})
	if err != nil {
		log.Fatal("failed to migrate database: ", err)
	}

	// 创建测试账户
	db.Create(&Account{Balance: 500.0}) // 账户A，ID=1
	db.Create(&Account{Balance: 300.0}) // 账户B，ID=2

	// 执行转账操作：从账户1向账户2转账100元
	err = TransferMoney(db, 1, 2, 100.0)
	if err != nil {
		fmt.Println("转账失败:", err)
	} else {
		fmt.Println("转账成功!")
	}

	// 显示转账后的账户余额
	var accountA, accountB Account
	db.First(&accountA, 1)
	db.First(&accountB, 2)
	fmt.Printf("账户A余额: %.2f\n", accountA.Balance)
	fmt.Printf("账户B余额: %.2f\n", accountB.Balance)

	// 显示交易记录
	var transactions []Transaction
	db.Find(&transactions)
	fmt.Println("交易记录:")
	for _, tx := range transactions {
		fmt.Printf("ID: %d, 从账户: %d, 到账户: %d, 金额: %.2f\n",
			tx.ID, tx.FromAccountID, tx.ToAccountID, tx.Amount)
	}

}

// TransferMoney 执行转账操作
func TransferMoney(db *gorm.DB, fromID, toID uint, amount float64) error {
	// 开始事务
	return db.Transaction(func(tx *gorm.DB) error {
		// 1. 检查转出账户是否存在且余额充足
		var fromAccount Account
		if err := tx.First(&fromAccount, fromID).Error; err != nil {
			return fmt.Errorf("转出账户不存在: %w", err)
		}

		if fromAccount.Balance < amount {
			return fmt.Errorf("转出账户余额不足，当前余额: %.2f", fromAccount.Balance)
		}

		// 2. 检查转入账户是否存在
		var toAccount Account
		if err := tx.First(&toAccount, toID).Error; err != nil {
			return fmt.Errorf("转入账户不存在: %w", err)
		}

		// 3. 更新转出账户余额
		if err := tx.Model(&fromAccount).Update("balance", fromAccount.Balance-amount).Error; err != nil {
			return fmt.Errorf("更新转出账户失败: %w", err)
		}

		// 4. 更新转入账户余额
		if err := tx.Model(&toAccount).Update("balance", toAccount.Balance+amount).Error; err != nil {
			return fmt.Errorf("更新转入账户失败: %w", err)
		}

		// 5. 记录交易
		transaction := Transaction{
			FromAccountID: fromID,
			ToAccountID:   toID,
			Amount:        amount,
		}
		if err := tx.Create(&transaction).Error; err != nil {
			return fmt.Errorf("记录交易失败: %w", err)
		}

		// 返回nil提交事务
		return nil
	})
}
