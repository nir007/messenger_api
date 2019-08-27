function createUser() {
  const form = $('#create_user_form');

  const data = {
    name: form.find('input[name="name"]').val(),
    uid: form.find('input[name="uid"]').val(),
    second: form.find('input[name="second"]').val(),
    gender: form.find('select[name="gender"]').val(),
    applicationID: form.find('input[name="appId"]').val(),
    email: form.find('input[name="email"]').val(),
    phone: form.find('input[name="phone"]').val(),
  };

  $.ajax({
    url: "/manage/users",
    method: "POST",
    dataType: 'json',
    data: JSON.stringify(data),
    headers: {
      "Content-Type":"application/json"
    },
    success: function (res) {

    },
    error: function (res) {
      alert('Error')
    }
  }).done(function (res) {
    console.log(res)
  })
}

function ajaxGetUsers(applicationId, beforeSend, success) {
  let params = beforeSend();
  console.log(params);
  $.ajax({
    url: "/manage/users?applicationid=" + applicationId + "&page=" + params.page + "&perpage=" + params.perPage,
    method: "GET",
    dataType: 'json',
    headers: {
      "Content-Type":"application/json"
    },
    success: function (res) {
      if (typeof success === "function") {
        success(res.result, res.total);
      }
    },
    error: function (res) {
      alert('Error')
    }
  }).done(function (res) {
    console.log(res)
  })
}

function createApplication() {
  const form = $('#create_application_form');

  form.find('input').keyup(function (e) {
    $(this).removeClass('is-invalid');
  });

  const data = {
    name: form.find('input[name="name"]').val(),
    description: form.find('input[name="description"]').val(),
    domains: form.find('input[name="domains"]').val(),
    managers: []
  };

  if (data.name.length === 0) {
    form.find('input[name="name"]').addClass('is-invalid');
    return
  }
  if (data.description.length === 0) {
    form.find('input[name="description"]').addClass('is-invalid');
    return
  }
  if (data.domains.length === 0) {
    form.find('input[name="domains"]').addClass('is-invalid');
    return
  }

  data.domains = data.domains.split(",");

  for (let i = 0; i < data.domains.length; i ++) {
    data.domains[i] = data.domains[i].replace(/ /g, '+');
  }

  $.ajax({
    url: "/manage/applications",
    method: "POST",
    dataType: 'json',
    data: JSON.stringify(data),
    headers: {
      "Content-Type":"application/json"
    },
    success: function (res) {
      if (typeof res.result.id !== 'undefined') {
        location.href = '/admin/application/' + res.result.id
      }
    },
    error: function (res) {
      alert('Error')
    }
  }).done(function (res) {
    console.log(res)
  })
}

function refreshSecretKey(appId) {
  let form = $('#update_application_form');
  $.ajax({
    url: "/manage/applications/secret-key",
    method: "PUT",
    dataType: 'json',
    data: JSON.stringify({id: appId}),
    headers: {
      "Content-Type":"application/json"
    },
    success: function (res) {
      form.find('input[name="secret"]').val(res.result)
    },
    error: function (res) {
      alert('Error')
    }
  }).done(function (res) {
    console.log(res)
  })
}

function getApplications(managerId, success) {
  $.ajax({
    url: "/manage/applications?manager_id=" + managerId,
    method: "GET",
    dataType: 'json',
    headers: {
      "Content-Type":"application/json"
    },
    success: function (res) {
      if (typeof success === "function") {
        success(res.result);
      }
    },
    error: function (res) {
      alert('Error')
    }
  }).done(function (res) {
    console.log(res)
  })
}

function deleteCookie( name ) {
  document.cookie = name + '=; expires=Thu, 01 Jan 1970 00:00:01 GMT;path=/';
}

function logout() {
  deleteCookie('jwt');
  deleteCookie('gopa');
  deleteCookie('jwtExpire');
  location.href = '/login'
}