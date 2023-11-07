from flask import Flask, request, render_template, redirect, url_for, session, jsonify
from flask_login import LoginManager, UserMixin, login_user, login_required, logout_user, current_user
from datetime import timedelta
from user_req import get_user_by_name

app = Flask(__name__)

# Set the secret key for session management
app.secret_key = "your_secret_key"
app.permanent_session_lifetime = timedelta(seconds=60)

# users = {
#     'user1': {'username': 'user1', 'password': 'password1'},
#     'user2': {'username': 'user2', 'password': 'password2'},
# }

class User(UserMixin):
    def __init__(self, id):
        self.id = id

login_manager = LoginManager()
login_manager.init_app(app)

@login_manager.user_loader
def load_user(user_id):
    return User(user_id)

@app.route('/')
def home():
    if 'username' in session:
        return f'Hello, {session["username"]}! <a href="/logout">Logout</a>'
    return 'You are not logged in. <a href="/login">Login</a>'

@app.route('/login', methods=['POST'])
def login():
    username = request.form['username']
    password = request.form['password']
    user = get_user_by_name(username)
    if isinstance(user, IndexError):
        return jsonify({"error": 'Invalid credentials'}), 400
    if username in user and user[-1] == password:
        user = User(username)
        login_user(user)
        session['username'] = username
        return redirect(url_for('home'))
    else:
        return jsonify({"error": 'Invalid credentials'}), 400
    return jsonify({"error": 'Invalid credentials'}), 400

@app.route('/logout')
@login_required
def logout():
    logout_user()
    session.pop('username', None)
    return redirect(url_for('home'))

@app.route('/check_user_logged_in', methods=['GET'])
def check_user_logged_in():
    print(request.headers["Cookie"])
    if current_user.is_authenticated:
        return jsonify({'logged_in': True, 'username': session['username']})
    else:
        return jsonify({'logged_in': False})

if __name__ == '__main__':
    app.run(debug=True, port=5002)