-- Единицы измерения.
CREATE TABLE units (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL DEFAULT ''
);

-- Типы продуктов.
CREATE TABLE product_types (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL DEFAULT ''
);

-- Товары.
CREATE TABLE items (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL DEFAULT '',
    price INTEGER NOT NULL DEFAULT 0,
    unit_id INTEGER NOT NULL REFERENCES units(id),
    product_type_id INTEGER NOT NULL REFERENCES product_types(id)
);

-- Статусы закаа.
CREATE TABLE order_statuses (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL DEFAULT ''
);

-- Роли пользователей.
CREATE TABLE roles (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL DEFAULT ''
);

-- Пользователи.
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL DEFAULT '',
    role_id INTEGER NOT NULL REFERENCES roles(id),
    login TEXT NOT NULL DEFAULT '',
    password TEXT NOT NULL DEFAULT ''
);

-- Типы оплат.
CREATE TABLE payment_types (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL DEFAULT ''
);

-- Заказы.
CREATE TABLE orders (
    id SERIAL PRIMARY KEY,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    design TEXT NOT NULL DEFAULT '',
    quantity INTEGER NOT NULL DEFAULT 0,
    status_id INTEGER NOT NULL REFERENCES order_statuses(id),
    price INTEGER NOT NULL DEFAULT 0,
    addr TEXT NOT NULL DEFAULT '',
    employee_id INTEGER NOT NULL REFERENCES users(id),
    deadline TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    payment_type_id INTEGER NOT NULL REFERENCES payment_types(id),
    delivery_info JSONB NOT NULL DEFAULT '{}'::JSONB,
    additional JSONB NOT NULL DEFAULT '{}'::JSONB,
);

-- Переходы между состояниями заказа.
CREATE TABLE order_transitions (
    id SERIAL PRIMARY KEY,
    prev_state_id INTEGER NOT NULL REFERENCES order_statuses(id),
    curr_state_id INTEGER NOT NULL REFERENCES order_statuses(id),
    transition_time TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- Заказы-товары.
CREATE TABLE orders_items (
    order_id INTEGER NOT NULL REFERENCES orders(id),
    item_id INTEGER NOT NULL REFERENCES items(id)
);

-- Типы заказчиков.
CREATE TABLE customer_types (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL DEFAULT ''
);

-- Заказчики.
CREATE TABLE customers (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL DEFAULT '',
    type_id INTEGER NOT NULL REFERENCES customer_types(id),
    addr TEXT NOT NULL DEFAULT '',
    email TEXT NOT NULL DEFAULT '',
    phone TEXT NOT NULL DEFAULT '',
    manager_name TEXT NOT NULL DEFAULT '',
    notes JSONB NOT NULL DEFAULT '{}'::JSONB
);

-- Промежуточная таблица заказчики - заказы.
CREATE TABLE customers_orders (
    customer_id INTEGER NOT NULL REFERENCES customers(id),
    order_id INTEGER NOT NULL REFERENCES orders(id)
);

-- Cостояния задачи.
CREATE TABLE task_states (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL DEFAULT ''
);

-- Задачи.
CREATE TABLE tasks (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL DEFAULT '',
    deadline TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    description TEXT NOT NULL DEFAULT '',
    employee_id INTEGER NOT NULL REFERENCES users(id),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    state_id INTEGER NOT NULL REFERENCES task_states(id)
);

-- Переходы между состояниями задачи.
CREATE TABLE task_transitions (
    id SERIAL PRIMARY KEY,
    prev_state_id INTEGER NOT NULL REFERENCES task_states(id),
    curr_state_id INTEGER NOT NULL REFERENCES task_states(id),
    transition_time TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- Поставщики.
CREATE TABLE suppliers (
    id SERIAL PRIMARY KEY,
    company_name TEXT NOT NULL DEFAULT '',
    manager_name TEXT NOT NULL DEFAULT '',
    email TEXT NOT NULL DEFAULT '',
    phone TEXT NOT NULL DEFAULT '',
    product_type_id INTEGER NOT NULL REFERENCES product_types(id),
    payment_type INTEGER NOT NULL REFERENCES payment_types(id),
    website_link TEXT NOT NULL DEFAULT '',
    comments TEXT NOT NULL DEFAULT ''
);

-- Поставки.
CREATE TABLE supplies (
    id SERIAL PRIMARY KEY,
    supplier_id INTEGER NOT NULL REFERENCES suppliers(id),
    supply_date TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    sum INTEGER NOT NULL DEFAULT 0,
    info JSONB NOT NULL DEFAULT '{}'::JSONB
);