<?php
session_start();
require "./database/db.php";

if (!isset($_SESSION['logged'])) {
  header('Location: /sansapp/php/home.php');
} else {
  $email = $_SESSION['email'];
  $query = "SELECT * FROM cuenta WHERE correo='$email'";
  $res = mysqli_query($conn, $query);
  $fila = mysqli_fetch_array($res);
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
  <form action="/sansapp/php/database/modify.php" method="post">
    <div class="container">
    <center> <h1> Modificar perfil </h1> </center>
      <?php
        if (isset($_SESSION['mensaje'])) {
          echo("<center><strong>" . $_SESSION['mensaje'] . "</strong></center><br>");
          unset($_SESSION['mensaje']);
        }
      ?>
      <label>💬 Nombre : </label>
      <input type="text" maxlength="50" placeholder="Introduce tu nombre" name="nombre" value="<?php echo($fila['nombre']); ?>" required>
      <label>📱 Telefono : </label>
      <input type="number" min="100000000" max="999999999" placeholder="Introduce tu telefono (ej: 912312678)" name="telef" value="<?php echo($fila['telefono']); ?>" required>
      <label>✉ Correo : </label>
      <input type="text" placeholder="Introduce correo nuevo" name="correo" value="<?php echo($fila['correo']); ?>" required>
      <label> <br> 🔑 Contraseña : </label>
      <input type="password" placeholder="Cambiar contraseña" name="pass" required>
      <input type="password" placeholder="Confirma tu contraseña" name="pass_conf" required>
      <button type="submit" name="env">✔ Modificar</button>

      <hr style="width:75%;">
      <center><button type="button" class="opposition"> <a href="mi_perfil.php">Atras</a></button> </center>
    </div>
  </form>
</body>
</html>
