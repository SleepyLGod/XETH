create table block
(
    id                        bigint       not null comment 'blockNum'
        primary key,
    created_at                bigint       not null comment 'timestamp',
    transaction_count         int          not null comment 'transactionCount',
    internal_transction_count int          null comment 'internaltransctioncount',
    miner_address             varchar(255) null comment 'mineraddress',
    block_reward              bigint       null,
    uncles_reward             varchar(255) null,
    total_difficulty          varchar(255) null,
    size                      int          null,
    gas_used                  int          null,
    gas_limit                 int          null,
    base_fee_per_gas          bigint       null,
    burnt_fees                bigint       null,
    extra_data                varchar(255) null,
    hash                      varchar(255) null,
    parent_hash               varchar(255) null,
    sha3_uncles               varchar(255) null,
    state_root                varchar(255) null,
    nounce                    varchar(255) null,
    difficulty                varchar(255) null
);

create table block_and_contract
(
    id          int unsigned auto_increment
        primary key,
    block_id    bigint       not null,
    contract_id int unsigned not null
);

create table block_and_internal_transaction
(
    id             int unsigned auto_increment
        primary key,
    block_num      bigint not null,
    transaction_id bigint not null
);

create table block_and_transaction
(
    id             int unsigned auto_increment
        primary key,
    block_id       bigint not null,
    transaction_id bigint not null
);

create table contract_info
(
    id                       int unsigned auto_increment
        primary key,
    created_at               bigint       not null,
    address                  varchar(255) null,
    created_block            bigint       null,
    created_transaction_hash varchar(255) null,
    creator_is_contract      tinyint(1)   null,
    create_value             bigint       null,
    creation_code            longtext     null,
    contract_code            longtext     null
);

create table erc20_transaction
(
    id               int unsigned auto_increment
        primary key,
    created_at       bigint       not null,
    block_num        bigint       null,
    transaction_hash varchar(255) null,
    token_address    varchar(255) null,
    `from`           varchar(255) null,
    `to`             varchar(255) null,
    from_is_contract tinyint(1)   null,
    to_is_contract   tinyint(1)   null,
    value            bigint       null
);

create table erc721_transaction
(
    id               int unsigned auto_increment
        primary key,
    created_at       bigint       not null,
    block_num        bigint       null,
    transaction_hash varchar(255) null,
    token_address    varchar(255) null,
    `from`           varchar(255) null,
    `to`             varchar(255) null,
    from_is_contract tinyint(1)   null,
    to_is_contract   tinyint(1)   null,
    token_id         int unsigned null
);

create table internal_transaction
(
    id                 int unsigned auto_increment
        primary key,
    created_at         bigint       not null,
    transaction_hash   varchar(255) null,
    type_trace_address varchar(255) null,
    `from`             varchar(255) null,
    `to`               varchar(255) null,
    from_is_contract   tinyint(1)   null,
    to_is_contract     tinyint(1)   null,
    value              bigint       null,
    calling_fuction    varchar(255) null,
    is_error           varchar(255) null
);

create table token_info
(
    id           int unsigned auto_increment
        primary key,
    address      varchar(255) null,
    name         varchar(255) null,
    symbol       varchar(255) null,
    total_supply varchar(255) null,
    `decimal`    int          null
);

create table transaction
(
    id                       int unsigned auto_increment
        primary key,
    block_num                bigint       null,
    transaction_hash         varchar(255) null,
    `from`                   varchar(255) null,
    `to`                     varchar(255) null,
    to_create                varchar(255) null,
    from_is_contract         tinyint(1)   null,
    to_is_contract           tinyint(1)   null,
    value                    bigint       null,
    gas_limit                int          null,
    gas_price                int          null,
    gas_used                 int          null,
    calling_function         varchar(255) null,
    is_error                 varchar(255) null,
    eip2718_type             varchar(255) null,
    base_fee_per_gas         bigint       null,
    max_fee_per_gas          bigint       null,
    max_priority_fee_per_gas bigint       null,
    created_at               bigint       not null
);

