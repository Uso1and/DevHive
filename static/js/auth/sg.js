document.addEventListener('DOMContentLoaded', function() {
            if (typeof particlesJS !== 'undefined') {
                particlesJS('particles-left', {
                    "particles": {
                        "number": {
                            "value": 40,
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
                            "value": 0.2,
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
                            "speed": 1,
                            "direction": "none",
                            "random": true,
                            "out_mode": "out"
                        }
                    }
                });
            }

            // Переключение видимости пароля
            document.querySelectorAll('.toggle-password').forEach(toggle => {
                toggle.addEventListener('click', function() {
                    const target = this.getAttribute('toggle-target');
                    const input = document.getElementById(target);
                    if (input) {
                        const type = input.getAttribute('type') === 'password' ? 'text' : 'password';
                        input.setAttribute('type', type);
                        this.classList.toggle('fa-eye-slash');
                    }
                });
            });

            // Обработка формы регистрации
            const signupForm = document.getElementById('signupForm');
            const errorMessageEl = document.getElementById('errorMessage');
            
            if (signupForm) {
                signupForm.addEventListener('submit', async function(e) {
                    e.preventDefault();

                    const username = document.getElementById('username').value;
                    const email = document.getElementById('email').value;
                    const password = document.getElementById('password').value;
                    const confirmPassword = document.getElementById('confirm_password').value;

                    // Простая проверка паролей
                    if (password !== confirmPassword) {
                        showErrorMessage('Пароли не совпадают');
                        return;
                    }

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
                            // Анимация успешной регистрации
                            signupForm.classList.add('success');
                            showSuccessMessage('Регистрация прошла успешно!');
                            
                            setTimeout(() => {
                                window.location.href = '/login';
                            }, 2000);
                        } else {
                            // Анимация ошибки
                            signupForm.classList.add('error');
                            setTimeout(() => signupForm.classList.remove('error'), 500);
                            
                            showErrorMessage(data.error || 'Ошибка регистрации');
                        }
                    } catch (err) {
                        showErrorMessage('Ошибка соединения с сервером');
                        console.error('Signup error:', err);
                    }
                });
            }

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