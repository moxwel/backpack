<?php
session_start();
require "./database/db.php";

if (!isset($_SESSION['logged'])) {
  header('Location: /sansapp/php/home.php');
  die();
}

if (!isset($_GET['id'])) {
  $_SESSION['mensaje'] = "Usa la opcion del menu.";
  header('Location: /sansapp/php/mi_perfil.php');
  die();
}

$id_cuenta = $_GET['id'];
$query = "SELECT * FROM cuenta WHERE id_cuenta='$id_cuenta'";
$res = mysqli_query($conn, $query);

$fila = mysqli_fetch_array($res);

if (is_null($fila)) {
  $_SESSION['mensaje'] = "No existe perfil.";
  header('Location: /sansapp/php/mi_perfil.php');
  die();
}

if (!($fila['id_cuenta'] == $_SESSION['id'])) {
  $_SESSION['mensaje'] = "Solo puedes borrar tu propia cuenta.";
  header('Location: /sansapp/php/mi_perfil.php');
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
</head>
<body>
  <center> <h1> <a href="/sansapp/php/home.php"> SansApp </a> </h1> </center>
    <div class="container">
    <center> <h1> Borrar cuenta </h1> </center>
      <center><h3>¿Quieres borrar tu cuenta?</h3></center>
      <center><button> <a href="mi_perfil.php">No</a></button> </center>
      <center><button style="background-color:red;"> <a href="/sansapp/php/database/deleteprofile.php?id=<?php echo($fila['id_cuenta']); ?>">❌ BORRAR</a></button> </center>
    </div>
</body>
</html>
