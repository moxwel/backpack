<?php
session_start();
require "./database/db.php";

if (!isset($_SESSION['logged'])) {
  header('Location: /sansapp/php/home.php');
  die();
}

if (!isset($_GET['id'])) {
  $_SESSION['mensaje'] = "Usa esta lista para seleccionar tus productos.";
  header('Location: /sansapp/php/mis_productos.php');
  die();
}

$id_prod = $_GET['id'];
$query = "SELECT * FROM producto WHERE id_producto='$id_prod'";
$res = mysqli_query($conn, $query);

$fila = mysqli_fetch_array($res);

if (is_null($fila)) {
  $_SESSION['mensaje'] = "No existe producto.";
  header('Location: /sansapp/php/mis_productos.php');
  die();
}

if (!($fila['vendedor'] == $_SESSION['id'])) {
  $_SESSION['mensaje'] = "No puedes borrar productos que no son tuyos.";
  header('Location: /sansapp/php/mis_productos.php');
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
    <center> <h1> Borrar producto </h1> </center>
      <center><h3>¿Quieres borrar el producto "<?php echo($fila['nombre']); ?>"?</h3></center>
      <center><button> <a href="mis_productos.php">No</a></button> </center>
      <center><button style="background-color:red;"> <a href="/sansapp/php/database/delete.php?id=<?php echo($fila['id_producto']); ?>">❌ BORRAR</a></button> </center>
    </div>
</body>
</html>
