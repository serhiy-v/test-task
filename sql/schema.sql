create table if not exists transactions
(
    transaction_id int not null,
    request_id int not null,
    terminal_id int not null,
    partner_object_id smallint not null,
    amount_total int not null,
    amount_original int not null,
    commission_ps float not null,
    commision_client int not null,
    commission_provider float not null,
    date_input timestamp not null,
    date_post timestamp not null,
    status text not null,
    payment_type text not null,
    payment_number text not null,
    service_id int not null,
    service text not null,
    payee_id bigint not null,
    payee_name text not null,
    payee_bank_mfo int not null,
    payee_bank_account text not null,
    payment_narrative text not null
);