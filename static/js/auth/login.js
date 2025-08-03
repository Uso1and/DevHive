// login.js
document.addEventListener('DOMContentLoaded', function() {
    const errorMessageEl = document.getElementById('errorMessage');
    const loginForm = document.getElementById('loginForm');
    const togglePassword = document.querySelector('.toggle-password');
    const passwordInput = document.getElementById('password');

    // Инициализация частиц
    if (typeof particlesJS !== 'undefined') {
        particlesJS('particles-js', {
            "particles": {
                "number": {
                    "value": 60,
                    "density": {
                        "enable": true,
                        "value_area": 800
                    }
                },
                "color": {
                    "value": "#58a6ff"
                },
                "shape": {
                    "type": "circle"
                },
                "opacity": {
                    "value": 0.3,
                    "random": true
                },
                "size": {
                    "value": 3,
                    "random": true
                },
                "line_linked": {
                    "enable": true,
                    "distance": 150,
                    "color": "#58a6ff",
                    "opacity": 0.1,
                    "width": 1
                },
                "move": {
                    "enable": true,
                    "speed": 1.5,
                    "direction": "none",
                    "random": true,
                    "out_mode": "out"
                }
            },
            "interactivity": {
                "detect_on": "canvas",
                "events": {
                    "onhover": {
                        "enable": true,
                        "mode": "grab"
                    },
                    "onclick": {
                        "enable": true,
                        "mode": "push"
                    }
                }
            }
        });
    }

    // Переключение видимости пароля
    togglePassword.addEventListener('click', function() {
        const type = passwordInput.getAttribute('type') === 'password' ? 'text' : 'password';
        passwordInput.setAttribute('type', type);
        this.classList.toggle('fa-eye-slash');
    });

    // Обработка формы входа
    loginForm.addEventListener('submit', async function(e) {
        e.preventDefault();

        const username = document.getElementById('username').value;
        const password = passwordInput.value;

        // Простая валидация
        if (!username || !password) {
            showErrorMessage('Пожалуйста, заполните все поля');
            return;
        }

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
                localStorage.setItem('token', data.token);
                
                // Добавляем анимацию успешного входа
                loginForm.classList.add('success');
                showSuccessMessage('Вход выполнен успешно!');
                
                setTimeout(() => {
                    window.location.href = '/main?token=' + data.token;
                }, 1000);
            } else {
                // Анимация ошибки
                loginForm.classList.add('error');
                setTimeout(() => loginForm.classList.remove('error'), 500);
                
                showErrorMessage(data.error || 'Ошибка входа');
            }
        } catch (err) {
            showErrorMessage('Ошибка соединения с сервером');
            console.error('Login error:', err);
        }
    });

    function showErrorMessage(message) {
        // Скрываем предыдущее сообщение перед показом нового
        if (errorMessageEl.timeout) {
            clearTimeout(errorMessageEl.timeout);
            errorMessageEl.classList.remove('show', 'success');
        }

        errorMessageEl.textContent = message;
        errorMessageEl.style.backgroundColor = '#f85149';
        errorMessageEl.classList.remove('success');
        
        // Запускаем анимацию появления
        setTimeout(() => {
            errorMessageEl.classList.add('show');
        }, 10);
        
        // Устанавливаем таймер для автоматического скрытия
        errorMessageEl.timeout = setTimeout(() => {
            errorMessageEl.classList.remove('show');
        }, 5000);
    }

    function showSuccessMessage(message) {
        // Скрываем предыдущее сообщение перед показом нового
        if (errorMessageEl.timeout) {
            clearTimeout(errorMessageEl.timeout);
            errorMessageEl.classList.remove('show');
        }

        errorMessageEl.textContent = message;
        errorMessageEl.style.backgroundColor = 'var(--success-color)';
        errorMessageEl.classList.add('success');
        
        // Запускаем анимацию появления
        setTimeout(() => {
            errorMessageEl.classList.add('show');
        }, 10);
        
        // Устанавливаем таймер для автоматического скрытия
        errorMessageEl.timeout = setTimeout(() => {
            errorMessageEl.classList.remove('show');
        }, 5000);
    }
});