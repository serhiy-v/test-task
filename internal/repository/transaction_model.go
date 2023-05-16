package repository

type Transaction struct {
	TransactionId      int     `csv:"TransactionId"`
	RequestId          int     `csv:"RequestId"`
	TerminalId         int     `csv:"TerminalId"`
	PartnerObjectId    int     `csv:"PartnerObjectId"`
	AmountTotal        int     `csv:"AmountTotal"`
	AmountOriginal     int     `csv:"AmountOriginal"`
	CommissionPs       float64 `csv:"CommissionPS"`
	CommissionClient   int     `csv:"CommissionClient"`
	CommissionProvider float64 `csv:"CommissionProvider"`
	DateInput          string  `csv:"DateInput"`
	DatePost           string  `csv:"DatePost"`
	Status             string  `csv:"Status"`
	PaymentType        string  `csv:"PaymentType"`
	PaymentNumber      string  `csv:"PaymentNumber"`
	ServiceId          int     `csv:"ServiceId"`
	Service            string  `csv:"Service"`
	PayeeId            int     `csv:"PayeeId"`
	PayeeName          string  `csv:"PayeeName"`
	PayeeBankMfo       int     `csv:"PayeeBankMfo"`
	PayeeBankAccount   string  `csv:"PayeeBankAccount"`
	PaymentNarrative   string  `csv:"PaymentNarrative"`
}
