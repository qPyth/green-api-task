$(document).ready(function() {
    $('#settings_button').click(function() {
        let idInstance = $('#id_instance').val();
        let apiTokenInstance = $('#apiToken_instance').val();
        let data = {
            idInstance: parseInt(idInstance),
            apiTokenInstance: apiTokenInstance
        };

        sendJsonReq(data, '/settings', 'POST')
    });

    $('#state_button').click(function() {
        let idInstance = $('#id_instance').val();
        let apiTokenInstance = $('#apiToken_instance').val();
        let data = {
            idInstance: parseInt(idInstance),
            apiTokenInstance: apiTokenInstance
        };

        sendJsonReq(data, '/state-instance', 'POST')
    });

    $('#send_message').click(function() {
        let idInstance = $('#id_instance').val();
        let apiTokenInstance = $('#apiToken_instance').val();
        let phoneNumber = $('#phone_number').val();
        let msg = $('#message').val();
        let data = {
            idInstance: parseInt(idInstance),
            apiTokenInstance: apiTokenInstance,
            phoneNumber: phoneNumber,
            message: msg
        };

        sendJsonReq(data, '/send-message', 'POST')
    });

    $('#send_file').click(function() {
        let idInstance = $('#id_instance').val();
        let apiTokenInstance = $('#apiToken_instance').val();
        let phoneNumber = $('#phone_number_file').val();
        let file = $('#file_link').val();
        let data = {
            idInstance: parseInt(idInstance),
            apiTokenInstance: apiTokenInstance,
            phoneNumber: phoneNumber,
            fileUrl: file
        };

        sendJsonReq(data, '/send-file', 'POST')
    });
});


function sendJsonReq(data, url, method) {
    $.ajax({
        url: url,
        type: method,
        contentType: 'application/json',
        data: JSON.stringify(data),
        success: function(response) {
            let jsonResponse = JSON.parse(response);
            let formattedJson = JSON.stringify(jsonResponse, null, 10);
            $('#result').text(formattedJson);
        },
        error: function(xhr, status, error) {
            var response = JSON.parse(xhr.responseText);
            if (response.hasOwnProperty('error')) {
                var errorMessage = response.error;
                $('#result').text(errorMessage);
            }
        }
    });
}