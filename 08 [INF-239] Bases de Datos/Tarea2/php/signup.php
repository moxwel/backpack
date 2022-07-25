<?php
session_start();
require "./database/db.php";

if (isset($_SESSION['logged']) and $_SESSION['logged']) {
  header('Location: /sansapp/php/home.php');
  die();
}
?>

<!DOCTYPE html>
<html>
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1">
  <link rel="stylesheet" href="./extras/logsign.css">
  <title> Regístrate | SansApp </title>
  <style>
     input[type=number] {
  width: 100%;
  margin: 8px 0;
  padding: 12px 20px;
  display: inline-block;
  border: 2px solid lightblue;
  box-sizing: border-box;
}

  </style>
</head>
<body>
  <center> <h1> <a href="/sansapp/php/home.php"> SansApp </a> </h1> </center>
  <form action="/sansapp/php/database/register.php" method="post">
    <div class="container">
    <center> <h1> Regístrate </h1> </center>
      <?php
        if (isset($_SESSION['mensaje'])) {
          echo("<center><strong>" . $_SESSION['mensaje'] . "</strong></center><br>");
          unset($_SESSION['mensaje']);
        }
      ?>
      <label>💬 Nombre : </label>
      <input type="text" maxlength="50" placeholder="Introduce tu nombre" name="nombre" required>
      <label>✉ Correo : </label>
      <input type="email" placeholder="Introduce tu correo" name="correo" required>
      <label>🏷 ROL USM : </label>
      <input type="number" min="1000000000" max="9999999999" placeholder="Introduce tu ROL USM (sin guion)" name="rol" required>
      <label>🎂 Fecha de nacimiento : </label>
      <input type="date" placeholder="Introduce tu fecha de nacimiento" name="fecha_nac" required>
      <label> <br> 🔑 Contraseña : </label>
      <input type="password" placeholder="Introduce tu contraseña" name="pass" required>
      <input type="password" placeholder="Confirma tu contraseña" name="pass_conf" required>
      <button type="submit" name="env">✔ Regístrate</button>

      <!--<input type="checkbox" checked="checked"> Remember me -->
      <hr style="width:75%;">
      <center><button type="button" class="opposition"> <a href="login.php">Ya tengo cuenta</a></button> </center>
      <!--Olvidaste tu <a href="#"> contraseña? </a>-->
    </div>
  </form>
</body>
</html>
