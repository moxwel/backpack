<?php
session_start();
require "./database/db.php";

if (!isset($_SESSION['logged'])) {
  header('Location: /sansapp/php/home.php');
  die();
}

$id_usuario = $_SESSION['id'];

if (!isset($_GET['id'])) {
  $_SESSION['mensaje'] = "Usa esta lista para seleccionar tus productos.";
  header('Location: /sansapp/php/carrito.php');
  die();
}

$id_prod = $_GET['id'];
$query = "SELECT productos_carrito.*, producto.nombre FROM productos_carrito JOIN producto ON productos_carrito.producto = producto.id_producto WHERE comprador='$id_usuario' AND producto='$id_prod'";
$res = mysqli_query($conn, $query);

$fila = mysqli_fetch_array($res);

if (is_null($fila)) {
  $_SESSION['mensaje'] = "No existe producto.";
  header('Location: /sansapp/php/carrito.php');
  die();
}

if (!($fila['comprador'] == $_SESSION['id'])) {
  $_SESSION['mensaje'] = "No puedes borrar productos que no son tuyos.";
  header('Location: /sansapp/php/carrito.php');
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
    <center> <h1> Quitar producto </h1> </center>
      <center><h3>¿Quieres quitar el producto "<?php echo($fila['nombre']); ?>" de tu carrito?</h3></center>
      <center><button> <a href="carrito.php">No</a></button> </center>
      <center><button style="background-color:red;"> <a href="/sansapp/php/database/deleteCart.php?id=<?php echo($fila['producto']); ?>">❌ BORRAR</a></button> </center>
    </div>
</body>
</html>
