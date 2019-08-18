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
      //location.href = '/admin/application/' + res.ID
    },
    error: function (res) {
      alert('Error')
    }
  }).done(function (res) {
    console.log(res)
  })
}