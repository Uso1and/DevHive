document.addEventListener('DOMContentLoaded', function() {
    // Общие элементы
    const errorMessageEl = document.getElementById('errorMessage');
    
    // Проверяем, на какой странице мы находимся
 if (document.getElementById('loginForm')) {
    const loginForm = document.getElementById('loginForm');
    
    loginForm.addEventListener('submit', async function(e) {
        e.preventDefault();
        
        const username = document.getElementById('username').value;
        const password = document.getElementById('password').value;
        
        try {
            const response = await fetch('/login', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify({ 
                    username: username,
                    password: password 
                }),
            });
            
            const data = await response.json();
            
            if (response.ok) {
                // Сохраняем токен и перенаправляем
                localStorage.setItem('token', data.token);
                window.location.href = '/profile';
            } else {
                errorMessageEl.textContent = data.error || 'Ошибка входа';
            }
        } catch (err) {
            errorMessageEl.textContent = 'Ошибка соединения с сервером';
            console.error('Login error:', err);
        }
    });
}
    
    if (document.getElementById('signupForm')) {
    const signupForm = document.getElementById('signupForm');
    
    signupForm.addEventListener('submit', async function(e) {
        e.preventDefault();
        
        const username = document.getElementById('username').value;
        const email = document.getElementById('email').value;
        const password = document.getElementById('password').value;
        
        try {
            const response = await fetch('/signup', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify({ 
                    username: username,
                    email: email,
                    password: password 
                }),
            });
            
            const data = await response.json();
            
            if (response.ok) {
                alert('Регистрация прошла успешно!');
                window.location.href = '/';
            } else {
                errorMessageEl.textContent = data.error || 'Ошибка регистрации';
            }
        } catch (err) {
            errorMessageEl.textContent = 'Ошибка соединения с сервером';
            console.error('Signup error:', err);
        }
    });
}
    
    if (document.getElementById('profileInfo')) {
        // Логика для страницы профиля
        const profileInfoEl = document.getElementById('profileInfo');
        const logoutBtn = document.getElementById('logoutBtn');
        
        // Загружаем данные профиля
        async function loadProfile() {
            try {
                const token = localStorage.getItem('token');
                
                // Получаем ID пользователя из токена (упрощённо)
                // В реальном приложении нужно декодировать JWT или получать ID другим способом
                const userId = 1; // Замени на реальный способ получения ID
                
                const response = await fetch(`/user/${userId}`, {
                    headers: {
                        'Authorization': `Bearer ${token}`,
                    },
                });
                
                if (response.ok) {
                    const user = await response.json();
                    profileInfoEl.innerHTML = `
                        <p><strong>ID:</strong> ${user.id}</p>
                        <p><strong>Имя пользователя:</strong> ${user.username}</p>
                        <p><strong>Email:</strong> ${user.email}</p>
                        <p><strong>Дата регистрации:</strong> ${new Date(user.created_at).toLocaleDateString()}</p>
                    `;
                } else {
                    profileInfoEl.innerHTML = '<p>Ошибка загрузки профиля</p>';
                }
            } catch (err) {
                console.error('Profile load error:', err);
                profileInfoEl.innerHTML = '<p>Ошибка соединения с сервером</p>';
            }
        }
        
        // Обработчик выхода
        logoutBtn.addEventListener('click', function() {
            localStorage.removeItem('token');
            window.location.href = '/';
        });
        
        // Загружаем профиль при загрузке страницы
        loadProfile();
    }
});