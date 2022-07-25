<?php
session_start();
require "./database/db.php";

if (isset($_SESSION['logged'])) {
  header('Location: /sansapp/php/home.php');
}
?>

<!DOCTYPE html>
<html>
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1">
  <link rel="stylesheet" href="./extras/logsign.css">
  <title> Iniciar Sesión | SansApp </title>
</head>
<body>
  <center> <h1> <a href="/sansapp/php/home.php"> SansApp </a> </h1> </center>
  <form action="/sansapp/php/database/log.php" method="post">
    <div class="container">
      <center> <h1> Iniciar Sesión </h1> </center>
      <?php
        if (isset($_SESSION['mensaje'])) {
          echo("<center><strong>" . $_SESSION['mensaje'] . "</strong></center><br>");
          unset($_SESSION['mensaje']);
        }
      ?>
      <label>✉ Correo : </label>
      <input type="email" placeholder="Introduce tu correo" name="correo" required>
      <label>🔑 Contraseña : </label>
      <input type="password" placeholder="Introduce tu contraseña" name="pass" required>
      <button type="submit" name="env">✔ Iniciar sesión</button>

      <!--<input type="checkbox" checked="checked"> Remember me -->
      <hr style="width:75%;">
      <center><button type="button" class="opposition"> <a href="signup.php">Crear cuenta nueva</a></button> </center>
      <!--Olvidaste tu <a href="#"> contraseña? </a>-->
    </div>
  </form>
</body>
</html>
