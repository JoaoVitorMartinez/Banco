package conta

import (
	".Banco/cliente"
)

type Conta struct {
	NumeroConta int
	Agencia     int
	saldo       float64
	Titular     cliente.Cliente
}

func (c *Conta) Sacar(valor float64) (bool, float64) {
	podeSacar := valor > 0 && c.saldo > valor

	if podeSacar {
		c.saldo -= valor
		return true, c.saldo
	}

	return false, c.saldo
}

func (c *Conta) Depositar(valor float64) (bool, float64) {
	podeDepositar := valor > 0

	if podeDepositar {
		c.saldo += valor
		return true, c.saldo
	}
	return false, c.saldo
}

func (c *Conta) Transferir(valor float64, cd *Conta) (bool, float64) {
	podeTrasferir, _ := c.Sacar(valor)

	if podeTrasferir {
		cd.Depositar(valor)
		return true, c.saldo
	}
	return false, c.saldo
}

func (c *Conta) Extrato() float64 {
	return c.saldo
}
