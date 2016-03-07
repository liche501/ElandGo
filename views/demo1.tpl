<!DOCTYPE html>
<html>
 
<head>
    <title>first example</title>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <script type="text/javascript" src="/views/js/jquery.js"></script>
    <script type="text/javascript" src="/views/js/avalon.js"></script>
    
    <script>
    $(function(){
		// $.post("/user/createuser",{username:"liche44444"}).done(function(data){
		// 	alert(data)
		// })
    })
        var vm = avalon.define({
            $id: "test",
            name: "司徒正美"
        })
    function fn_redirect(){
        $.post("/demo1","").done(function(){

        })
    }
    </script>
</head>
 
<body ms-controller="test">
    <input ms-duplex="name">
    <p>Hello,{{name}}!</p>
	<p>template,[[.username]]</p>
    <form action="/demo1" method="post">
        <input type="submit" value="submit">
    </form>
    <input type="button" value="redirect" onclick="fn_redirect()">
</body>
 
</html>
