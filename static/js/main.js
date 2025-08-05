document.addEventListener('DOMContentLoaded', function() {
  const logoutBtn = document.getElementById('logoutBtn');
  const profileLink = document.getElementById('profileLink');
  const searchfrLink = document.getElementById('searchfrLink');
  const createDiscussionForm = document.getElementById('createDiscussionForm');
  const discussionsContainer = document.getElementById('discussionsContainer');
  const discussionMessage = document.getElementById('discussionMessage');
  
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

  searchfrLink.addEventListener('click', function(e){
    e.preventDefault();
    const token = localStorage.getItem('token');
    if (!token){
      window.location.href = '/';
      return;
    }
    window.location.href = `/searchfr?token=${encodeURIComponent(token)}`;
  });

  // Обработчик создания обсуждения
   if (createDiscussionForm) {
    createDiscussionForm.addEventListener('submit', async function(e) {
      e.preventDefault();
      
      const title = document.getElementById('discussionTitle').value;
      const description = document.getElementById('discussionDescription').value;
      const token = localStorage.getItem('token');
      
      if (!title || !description) {
        showMessage('Заполните все поля', 'error');
        return;
      }
      
      try {
        const response = await fetch('/discussions', {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json',
            'Authorization': `Bearer ${token}`
          },
          body: JSON.stringify({
            title: title,
            description: description
          })
        });
        
        if (!response.ok) {
          const errorData = await response.json();
          throw new Error(errorData.error || 'Ошибка при создании обсуждения');
        }
        
        const data = await response.json();
        showMessage('Обсуждение успешно создано!', 'success');
        createDiscussionForm.reset();
        loadDiscussions();
      } catch (err) {
        showMessage(err.message || 'Ошибка соединения с сервером', 'error');
        console.error('Create discussion error:', err);
      }
    });
  }
  
  // Функция для отображения сообщений
  function showMessage(text, type) {
    discussionMessage.textContent = text;
    discussionMessage.className = `message ${type}`;
    discussionMessage.style.display = 'block';
    
    setTimeout(() => {
      discussionMessage.style.display = 'none';
    }, 5000);
  }
  
  // Функция загрузки обсуждений
  async function loadDiscussions() {
    try {
      const response = await fetch('/discussions', {
        headers: {
          'Authorization': `Bearer ${token}`
        }
      });
      
      if (response.ok) {
        const discussions = await response.json();
        renderDiscussions(discussions);
      } else {
        console.error('Failed to load discussions');
      }
    } catch (err) {
      console.error('Error loading discussions:', err);
    }
  }
  
  // Функция отображения обсуждений
  function renderDiscussions(discussions) {
    discussionsContainer.innerHTML = '';
    
    if (discussions.length === 0) {
      discussionsContainer.innerHTML = '<p>Нет обсуждений</p>';
      return;
    }
    
    discussions.forEach(discussion => {
      const discussionElement = document.createElement('div');
      discussionElement.className = 'discussion';
      discussionElement.innerHTML = `
        <h3>${discussion.title}</h3>
        <p>${discussion.description}</p>
        <div class="discussion-meta">
          <span>Автор: ${discussion.creator_id}</span>
          <span>Дата: ${new Date(discussion.created_at).toLocaleString()}</span>
        </div>
      `;
      discussionsContainer.appendChild(discussionElement);
    });
  }
  
  // Загружаем обсуждения при загрузке страницы
  loadDiscussions();
});