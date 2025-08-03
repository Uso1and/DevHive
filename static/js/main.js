document.addEventListener('DOMContentLoaded', function() {
  const logoutBtn = document.getElementById('logoutBtn');
  const profileLink = document.getElementById('profileLink');

  // Проверка токена при загрузке
  const token = localStorage.getItem('token');
  if (!token) {
    window.location.href = '/';
    return;
  }

  // Обработчик выхода
  logoutBtn.addEventListener('click', function() {
    localStorage.removeItem('token');
    window.location.href = '/';
  });

  // Обработчик перехода в профиль
  profileLink.addEventListener('click', function(e) {
    e.preventDefault();
    const token = localStorage.getItem('token');
    if (!token) {
      window.location.href = '/';
      return;
    }
    window.location.href = `/profile?token=${encodeURIComponent(token)}`;
  });
});