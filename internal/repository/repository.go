package repository

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

type Repository struct {
	Conn *sql.DB
}

func NewRepository(conn *sql.DB) *Repository {
	return &Repository{Conn: conn}
}

func NewConnection(dsn string) (*sql.DB, error) {
	conn, err := sql.Open("postgres", dsn)

	if err != nil {
		return nil, err
	}

	pingErr := conn.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	return conn, err
}

func (r *Repository) AddTransaction(transaction *Transaction) error {
	q := "INSERT INTO transactions (transaction_id, request_id,terminal_id,partner_object_id,amount_total," +
		"amount_original,commission_ps,commision_client,commission_provider,date_input,date_post,status,payment_type," +
		"payment_number,service_id,service,payee_id,payee_name,payee_bank_mfo,payee_bank_account, payment_narrative) " +
		"VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15,$16,$17,$18,$19,$20,$21)"
	_, err := r.Conn.Exec(q, transaction.TransactionId, transaction.RequestId, transaction.TerminalId,
		transaction.PartnerObjectId, transaction.AmountTotal, transaction.AmountOriginal, transaction.CommissionPs,
		transaction.CommissionClient, transaction.CommisionProvider, transaction.DateInput, transaction.DatePost,
		transaction.Status, transaction.PaymentType, transaction.PaymentNumber, transaction.ServiceId, transaction.Service,
		transaction.PayeeId, transaction.PayeeName, transaction.PayeeBankMfo, transaction.PayeeBankAccount, transaction.PaymentNarrative)

	return err
}

func (r *Repository) GetByTransactionID(id string) (*Transaction, error) {
	q := "SELECT transaction_id, request_id,terminal_id,partner_object_id,amount_total," +
		"amount_original,commission_ps,commision_client,commission_provider,date_input,date_post,status,payment_type," +
		"payment_number,service_id,service,payee_id,payee_name,payee_bank_mfo,payee_bank_account, payment_narrative " +
		"FROM transactions WHERE transaction_id=$1"
	transaction := new(Transaction)

	err := r.Conn.QueryRow(q, id).Scan(&transaction.TransactionId, &transaction.RequestId, &transaction.TerminalId,
		&transaction.PartnerObjectId, &transaction.AmountTotal, &transaction.AmountOriginal, &transaction.CommissionPs,
		&transaction.CommissionClient, &transaction.CommisionProvider, &transaction.DateInput, &transaction.DatePost,
		&transaction.Status, &transaction.PaymentType, &transaction.PaymentNumber, &transaction.ServiceId, &transaction.Service,
		&transaction.PayeeId, &transaction.PayeeName, &transaction.PayeeBankMfo, &transaction.PayeeBankAccount, &transaction.PaymentNarrative)

	if err != nil {
		return nil, err
	}

	return transaction, nil
}

func (r *Repository) GetByTerminalID(id string) (*Transaction, error) {
	q := "SELECT transaction_id, request_id,terminal_id,partner_object_id,amount_total," +
		"amount_original,commission_ps,commision_client,commission_provider,date_input,date_post,status,payment_type," +
		"payment_number,service_id,service,payee_id,payee_name,payee_bank_mfo,payee_bank_account, payment_narrative " +
		"FROM transactions WHERE terminal_id=$1"
	transaction := new(Transaction)

	err := r.Conn.QueryRow(q, id).Scan(&transaction.TransactionId, &transaction.RequestId, &transaction.TerminalId,
		&transaction.PartnerObjectId, &transaction.AmountTotal, &transaction.AmountOriginal, &transaction.CommissionPs,
		&transaction.CommissionClient, &transaction.CommisionProvider, &transaction.DateInput, &transaction.DatePost,
		&transaction.Status, &transaction.PaymentType, &transaction.PaymentNumber, &transaction.ServiceId, &transaction.Service,
		&transaction.PayeeId, &transaction.PayeeName, &transaction.PayeeBankMfo, &transaction.PayeeBankAccount, &transaction.PaymentNarrative)

	if err != nil {
		return nil, err
	}

	return transaction, nil
}

func (r *Repository) GetByStatus(status string) ([]*Transaction, error) {
	q := "SELECT * FROM transactions WHERE status=$1"

	rows, err := r.Conn.Query(q, status)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	transactions := make([]*Transaction, 0)

	for rows.Next() {
		transaction := new(Transaction)
		err := rows.Scan(&transaction.TransactionId, &transaction.RequestId, &transaction.TerminalId,
			&transaction.PartnerObjectId, &transaction.AmountTotal, &transaction.AmountOriginal, &transaction.CommissionPs,
			&transaction.CommissionClient, &transaction.CommisionProvider, &transaction.DateInput, &transaction.DatePost,
			&transaction.Status, &transaction.PaymentType, &transaction.PaymentNumber, &transaction.ServiceId, &transaction.Service,
			&transaction.PayeeId, &transaction.PayeeName, &transaction.PayeeBankMfo, &transaction.PayeeBankAccount, &transaction.PaymentNarrative)

		if err != nil {
			return nil, err
		}

		transactions = append(transactions, transaction)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return transactions, nil
}

func (r *Repository) GetByPaymentType(payment string) ([]*Transaction, error) {
	q := "SELECT * FROM transactions WHERE payment_type=$1"

	rows, err := r.Conn.Query(q, payment)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	transactions := make([]*Transaction, 0)

	for rows.Next() {
		transaction := new(Transaction)
		err := rows.Scan(&transaction.TransactionId, &transaction.RequestId, &transaction.TerminalId,
			&transaction.PartnerObjectId, &transaction.AmountTotal, &transaction.AmountOriginal, &transaction.CommissionPs,
			&transaction.CommissionClient, &transaction.CommisionProvider, &transaction.DateInput, &transaction.DatePost,
			&transaction.Status, &transaction.PaymentType, &transaction.PaymentNumber, &transaction.ServiceId, &transaction.Service,
			&transaction.PayeeId, &transaction.PayeeName, &transaction.PayeeBankMfo, &transaction.PayeeBankAccount, &transaction.PaymentNarrative)

		if err != nil {
			return nil, err
		}

		transactions = append(transactions, transaction)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return transactions, nil
}

func (r *Repository) GetByPeriod(day Date) ([]*Transaction, error) {
	q := "SELECT * FROM transactions WHERE date_post >= $1 AND date_post <= $2"

	rows, err := r.Conn.Query(q, day.DateFrom, day.DateTo)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	transactions := make([]*Transaction, 0)

	for rows.Next() {
		transaction := new(Transaction)
		err := rows.Scan(&transaction.TransactionId, &transaction.RequestId, &transaction.TerminalId,
			&transaction.PartnerObjectId, &transaction.AmountTotal, &transaction.AmountOriginal, &transaction.CommissionPs,
			&transaction.CommissionClient, &transaction.CommisionProvider, &transaction.DateInput, &transaction.DatePost,
			&transaction.Status, &transaction.PaymentType, &transaction.PaymentNumber, &transaction.ServiceId, &transaction.Service,
			&transaction.PayeeId, &transaction.PayeeName, &transaction.PayeeBankMfo, &transaction.PayeeBankAccount, &transaction.PaymentNarrative)

		if err != nil {
			return nil, err
		}

		transactions = append(transactions, transaction)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return transactions, nil
}

func (r *Repository) GetByNarrative(text string) ([]*Transaction, error) {
	q := "SELECT * FROM transactions WHERE payment_narrative LIKE " +
		"'%" + text + "%'"

	rows, err := r.Conn.Query(q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	transactions := make([]*Transaction, 0)

	for rows.Next() {
		transaction := new(Transaction)
		err := rows.Scan(&transaction.TransactionId, &transaction.RequestId, &transaction.TerminalId,
			&transaction.PartnerObjectId, &transaction.AmountTotal, &transaction.AmountOriginal, &transaction.CommissionPs,
			&transaction.CommissionClient, &transaction.CommisionProvider, &transaction.DateInput, &transaction.DatePost,
			&transaction.Status, &transaction.PaymentType, &transaction.PaymentNumber, &transaction.ServiceId, &transaction.Service,
			&transaction.PayeeId, &transaction.PayeeName, &transaction.PayeeBankMfo, &transaction.PayeeBankAccount, &transaction.PaymentNarrative)

		if err != nil {
			return nil, err
		}

		transactions = append(transactions, transaction)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return transactions, nil
}
