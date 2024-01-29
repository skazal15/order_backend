import psycopg2
import pandas as pd

class PostgresDatabase:
    def __init__(self, db_params):
        self.conn = psycopg2.connect(**db_params)
        self.conn.autocommit = True
        self.cursor = self.conn.cursor()

    def create_table(self, sql_command):
        self.cursor.execute(sql_command)

    def insert_data_from_csv(self, file_path, table_name):
        df = pd.read_csv(file_path)
        df.fillna(0, inplace=True)
        for _, row in df.iterrows():
            placeholders = ', '.join(['%s'] * len(row))
            sql = f"INSERT INTO {table_name} VALUES ({placeholders})"
            self.cursor.execute(sql, tuple(row))

    def close(self):
        self.cursor.close()
        self.conn.close()

if __name__ == "__main__":
    # Database connection parameters6
    """edit variables with your local database"""
    db_params = {
        'database': 'order',
        'user': 'root',
        'password': 'root',
        'host': 'localhost',
        'port': 5432
    }

    db = PostgresDatabase(db_params)

    # SQL commands for creating tables, adjusted for correct order
    customer_companies_cmd = '''
        CREATE TABLE IF NOT EXISTS customer_companies (
            company_id INTEGER PRIMARY KEY,
            company_name TEXT NOT NULL
        );
    '''
    customers_cmd = '''
        CREATE TABLE IF NOT EXISTS customers (
            user_id TEXT PRIMARY KEY,
            login TEXT,
            password TEXT,
            name TEXT,
            company_id INTEGER REFERENCES customer_companies(company_id),
            credit_cards TEXT
        );
    '''
    deliveries_cmd = '''
        CREATE TABLE IF NOT EXISTS deliveries (
            id INTEGER PRIMARY KEY,
            order_item_id INTEGER REFERENCES order_items(id),
            delivered_quantity INTEGER
        );
    '''
    orders_cmd = '''
        CREATE TABLE IF NOT EXISTS orders (
            id INTEGER PRIMARY KEY,
            created_at TIMESTAMP,
            order_name TEXT,
            customer_id TEXT REFERENCES customers(user_id)
        );
    '''
    order_items_cmd = '''
        CREATE TABLE IF NOT EXISTS order_items (
            id INTEGER PRIMARY KEY,
            order_id INTEGER REFERENCES orders(id),
            price_per_unit NUMERIC,
            quantity INTEGER,
            product TEXT
        );
    '''
    deliveries_cmd = '''
        CREATE TABLE IF NOT EXISTS deliveries (
            id INTEGER PRIMARY KEY,
            order_item_id INTEGER REFERENCES order_items(id),
            delivered_quantity INTEGER
        );
    '''

    # Create tables
    db.create_table(customer_companies_cmd)
    db.create_table(customers_cmd)
    db.create_table(orders_cmd)
    db.create_table(order_items_cmd)
    db.create_table(deliveries_cmd)

    # Insert data from CSV files
    """edit path csv with your local path"""
    db.insert_data_from_csv('Test task - Postgres - customer_companies.csv', 'customer_companies')
    db.insert_data_from_csv('Test task - Postgres - customers.csv', 'customers')
    db.insert_data_from_csv('Test task - Postgres - orders.csv', 'orders')
    db.insert_data_from_csv('Test task - Postgres - order_items.csv', 'order_items')
    db.insert_data_from_csv('Test task - Postgres - deliveries.csv', 'deliveries')

    # Close the database connection
    db.close()
