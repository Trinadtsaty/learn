package bank

type account struct { // неэкспортируемый
    balance float64
}

// Экспортируемый метод
func (a *account) Deposit(amount float64) {
    a.balance += amount
}

// Фабрика
func NewAccount() *account {
    return &account{balance: 0}
}

func (a *account) GetBalance() float64 { return a.balance }