import psycopg2
from config import config

def get_user_by_name(username):
    conn = None
    try:
        command = f'''
        SELECT * FROM users
        WHERE username = '{username}';        
        '''
        params = config(section="authdb")
        # connect to the PostgreSQL server
        conn = psycopg2.connect(**params)
        cur = conn.cursor()

        cur.execute(command)
        books = cur.fetchall()[0]
        # close communication with the PostgreSQL database server
        cur.close()
        # commit the changes
        conn.commit()
    except (Exception, psycopg2.DatabaseError) as error:
        return error
    finally:
        if conn is not None:
            conn.close()

    return books

# print(get_user_by_name("user1"))