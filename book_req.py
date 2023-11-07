import psycopg2
from config import config

def get_books():
    conn = None
    try:
        command = '''
        SELECT * FROM books;        
        '''
        params = config()
        # connect to the PostgreSQL server
        conn = psycopg2.connect(**params)
        cur = conn.cursor()

        cur.execute(command)
        books = cur.fetchall()
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

def get_authors():
    conn = None
    try:
        command = '''
        SELECT * FROM authors;        
        '''
        params = config()
        conn = psycopg2.connect(**params)
        cur = conn.cursor()
        cur.execute(command)
        books = cur.fetchall()
        cur.close()
        conn.commit()
    except (Exception, psycopg2.DatabaseError) as error:
        return error
    finally:
        if conn is not None:
            conn.close()
    return books

def delete_author_by_id(id):
    conn = None
    try:
        command = f'''
        DELETE FROM authors
        WHERE author_id = {id}        
        '''
        params = config()
        # connect to the PostgreSQL server
        conn = psycopg2.connect(**params)
        cur = conn.cursor()

        cur.execute(command)
        # close communication with the PostgreSQL database server
        cur.close()
        # commit the changes
        conn.commit()
    except (Exception, psycopg2.DatabaseError) as error:
        return error
    finally:
        if conn is not None:
            conn.close()

def get_book_by_title(title):
    conn = None
    try:
        command = '''
        SELECT *
        FROM books
        WHERE LOWER(book_title) = '{title}'
        '''.format(title=title.lower())
        params = config()
        conn = psycopg2.connect(**params)
        cur = conn.cursor()
        cur.execute(command)
        books = cur.fetchall()
        cur.close()
        conn.commit()
    except (Exception, psycopg2.DatabaseError) as error:
        return error
    finally:
        if conn is not None:
            conn.close()
    return books

def get_book_by_author(author):
    conn = None
    try:
        command = '''
        WITH select_fullname as
        (
            SELECT author_id, LOWER(author_name || ' ' || author_surname) as author_fullname
            FROM authors
        )
        
        SELECT * FROM books as a
        INNER JOIN select_fullname as b
        ON a.author_id = b.author_id
        WHERE b.author_fullname = '{author}'
        '''.format(author=author.lower())
        params = config()
        conn = psycopg2.connect(**params)
        cur = conn.cursor()
        cur.execute(command)
        books = cur.fetchall()
        cur.close()
        conn.commit()
    except (Exception, psycopg2.DatabaseError) as error:
        return error
    finally:
        if conn is not None:
            conn.close()
    return books

def get_books_by_id(id):
    conn = None
    try:
        command = '''
        SELECT * FROM books
        WHERE book_id = {id};        
        '''.format(id=id)
        params = config()
        # connect to the PostgreSQL server
        conn = psycopg2.connect(**params)
        cur = conn.cursor()

        cur.execute(command)
        books = cur.fetchall()
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

def insert_author(first_name, last_name):
    # authors = get_authors()
    # print(authors)
    conn = None
    try:
        check_presence = f'''
        SELECT author_id FROM authors WHERE author_name = '{first_name}'
        AND author_surname = '{last_name}'
        LIMIT 1;
        '''
        insert_author = f'''
        INSERT INTO authors (author_name, author_surname)
        VALUES ('{first_name}', '{last_name}')
        RETURNING author_id;
        '''
        params = config()
        # connect to the PostgreSQL server
        conn = psycopg2.connect(**params)
        cur = conn.cursor()

        cur.execute(check_presence)
        id = cur.fetchall()

        if not id:
            cur.execute(insert_author)
            id = cur.fetchall()


        # close communication with the PostgreSQL database server
        cur.close()
        # commit the changes
        conn.commit()
    except (Exception, psycopg2.DatabaseError) as error:
        return error
    finally:
        if conn is not None:
            conn.close()

    return id[0][0]

def insert_book(title, genre, year, author_id, file_path):
    # authors = get_authors()
    # print(authors)
    conn = None
    try:
        insert_book = f'''
        INSERT INTO books (book_title, book_year, book_genre, author_id, file_path)
        VALUES ('{title}', {year}, '{genre}', {author_id}, '{file_path}')
        '''
        params = config()
        # connect to the PostgreSQL server
        conn = psycopg2.connect(**params)
        cur = conn.cursor()

        cur.execute(insert_book)
        # close communication with the PostgreSQL database server
        cur.close()
        # commit the changes
        conn.commit()
    except (Exception, psycopg2.DatabaseError) as error:
        return error
    finally:
        if conn is not None:
            conn.close()

print(insert_book("Totaler Negger Tod", "Politics", 1940, 3, "kek"))
print(get_books())