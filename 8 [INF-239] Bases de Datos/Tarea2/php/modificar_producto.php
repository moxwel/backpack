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
  $_SESSION['mensaje'] = "No puedes modificar productos que no son tuyos.";
  header('Location: /sansapp/php/mis_productos.php');
  die();
}

$_SESSION['temp_id_producto'] = $id_prod;
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

textarea {
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

  <form action="/sansapp/php/database/modify_prod.php" method="post">

    <div class="container">
    <center> <h1> Modificar producto </h1> </center>


      <?php
        if (isset($_SESSION['mensaje'])) {
          echo("<center><strong>" . $_SESSION['mensaje'] . "</strong></center><br>");
          unset($_SESSION['mensaje']);
        }
      ?>
      <label>💬 Nombre del producto: </label>
      <input type="text" maxlength="20" placeholder="Introduce el nombre del producto" name="nombre" value="<?php echo($fila['nombre']); ?>" required>
      <label>💲 Precio : </label>
      <input type="number" min="0" max="9999999" placeholder="Introduce precio" name="precio" value="<?php echo($fila['precio']); ?>" required>
      <label>🏷 Etiqueta : </label>
      <input type="text" maxlength="20" placeholder="Introduce etiqueta" name="etiqueta" value="<?php echo($fila['etiqueta']); ?>" required>
      <label>📄 Descripcion : </label><br>
      <textarea name="descripcion" rows="20" cols="200"placeholder="Introduce una descripcion del producto" required><?php echo($fila['descripcion']); ?></textarea>
      <button type="submit" name="env">✔ Modificar</button>

      <hr style="width:75%;">
      <center><button type="button" class="opposition"> <a href="mis_productos.php">Atras</a></button> </center>
    </div>
  </form>
</body>
</html>
