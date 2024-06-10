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
    unit_id INTEGER NOT NULL REFERENCES units(id)
);

-- Статусы заказа.
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
    created_at INTEGER NOT NULL DEFAULT 0,
    design TEXT NOT NULL DEFAULT '',
    quantity INTEGER NOT NULL DEFAULT 0,
    status_id INTEGER NOT NULL REFERENCES order_statuses(id),
    price INTEGER NOT NULL DEFAULT 0,
    addr TEXT NOT NULL DEFAULT '',
    employee_id INTEGER NOT NULL REFERENCES users(id),
    deadline INTEGER NOT NULL DEFAULT 0,
    payment_type_id INTEGER NOT NULL REFERENCES payment_types(id),
    delivery_info JSONB NOT NULL DEFAULT '{}'::JSONB,
    additional JSONB NOT NULL DEFAULT '{}'::JSONB
);

-- Переходы между состояниями заказа.
CREATE TABLE order_transitions (
    id SERIAL PRIMARY KEY,
    prev_state_id INTEGER NOT NULL REFERENCES order_statuses(id),
    curr_state_id INTEGER NOT NULL REFERENCES order_statuses(id),
    transition_time INTEGER NOT NULL DEFAULT 0
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
    notes TEXT NOT NULL DEFAULT ''
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
    deadline INTEGER NOT NULL DEFAULT 0,
    description TEXT NOT NULL DEFAULT '',
    employee_id INTEGER NOT NULL REFERENCES users(id),
    created_at INTEGER NOT NULL DEFAULT 0,
    state_id INTEGER NOT NULL REFERENCES task_states(id)
);

-- Переходы между состояниями задачи.
CREATE TABLE task_transitions (
    id SERIAL PRIMARY KEY,
    task_id INTEGER NOT NULL REFERENCES tasks(id), -- Задача.
    prev_state_id INTEGER NOT NULL REFERENCES task_states(id),
    curr_state_id INTEGER NOT NULL REFERENCES task_states(id),
    transition_time INTEGER NOT NULL DEFAULT 0
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

