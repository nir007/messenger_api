function registration() {
  const form = $('#registration_form');
  const data = {
    name: form.find('input[name="name"]').val(),
    second: form.find('input[name="second"]').val(),
    email: form.find('input[name="email"]').val(),
    password: form.find('input[name="password"]').val()
  };

  $.ajax({
    url: "/registration",
    method: "POST",
    dataType: 'json',
    data: JSON.stringify(data),
    headers: {
      "Content-Type":"application/json"
    },
    success: function (res) {
      location.href = '/login'
    },
    error: function (res) {
      alert('Error')
    }
  }).done(function (res) {
    console.log(res)
  })
}

function login() {
  const form = $('#login_form');
  const data = {
    email: form.find('input[name="email"]').val(),
    password: form.find('input[name="password"]').val()
  };

  $.ajax({
    url: "/login",
    method: "POST",
    dataType: 'json',
    data: JSON.stringify(data),
    headers: {
      "Content-Type":"application/json"
    },
    success: function (res) {
      document.cookie = "jwt="+res.token; // обновляем только куки с именем 'user'
      document.cookie = "jwtExpire="+res.expire; // обновляем только куки с именем 'user'
      location.href = '/admin/dashboard'
    },
    error: function (res) {
      alert('Error')
    }
  }).done(function (res) {
    console.log(res)
  })
}