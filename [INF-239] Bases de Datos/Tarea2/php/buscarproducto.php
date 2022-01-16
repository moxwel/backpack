<?php
session_start();
require "./database/db.php";

if (!isset($_GET['q'])) {
  header('Location: /sansapp/php/home.php');
}

$busqueda = $_GET['q'];
$query = "SELECT * FROM producto WHERE nombre LIKE '%$busqueda%'";
$res = mysqli_query($conn, $query);
?>

<!DOCTYPE html>
<html>
<head>
<meta charset="UTF-8">
<meta name="viewport" content="width=device-width, initial-scale=1">
<title> SansApp </title>
<style>
body {
  margin: 0;
  font-family: Arial, Helvetica, sans-serif;
  background-color: #4BC3C7;
}

.topnav {
  overflow: hidden;
  background-color: lightblue;
}

.topnav a {
  float: left;
  color: black;
  text-align: center;
  padding: 14px 16px;
  text-decoration: none;
  font-size: 17px;
}

.topnav a:hover {
  background-color: #E7A12F;
  color: black;
}

.topnav a.active {
  background-color: #E7A12F;
  color: black;

}
.topnav a.active:hover {
  background-color: white;
  color: black;

}

.topnav b {
  float: right;
  color: black;
  text-align: center;
  padding: 14px 16px;
  text-decoration: none;
  font-size: 17px;
}
.topnav b:hover {
  background-color: #E7A12F;
  color: black;
}

.topnav c {
  float: left;
  color: black;
  text-align: match-parent ;
  padding: 14px 16px;
  text-decoration: none;
  font-size: 17px;
}
.topnav d {
  float: left;
  color: black;
  text-align: match-parent ;
  padding: 14px 0px;
  text-decoration: none;
  font-size: 17px;
}



.column {
  float: left;
  width: 60%;
  padding: 10px;

}


.row:after {
  content: "";
  display: table;
  clear: both;
}



table {
  font-family: arial, sans-serif;
  border-collapse: collapse;
  width: 100%;
}

td, th {
  border: 1px solid #dddddd;
  text-align: left;
  padding: 8px;
}

tr:nth-child(even) {
  background-color: #dddddd;
}

</style>
</head>
<body>

<?php require "./extras/topnav.php" ?>
<h2>Buscar productos</h2>
<h3>" <?php if ($_GET['q'] !== ''){echo($_GET['q']);}else{echo("Todos los productos");} ?> "</h3>
<div class="row">
    <div class="column" style="background-color:lightblue;">
      <h2>Productos encontrados</h2>
      <table>
        <tr>
          <th>Nombre producto</th>
          <th>Precio</th>
          <th>Cantidad</th>
          <th>Ventas</th>
          <th>Calificacion</th>

        </tr>


        <?php

        $fila = mysqli_fetch_array($res);
        while (!is_null($fila)) {
        ?>

        <tr>
          <td><a href="producto.php?id=<?php echo($fila['id_producto']); ?>"><?php echo($fila['nombre']); ?></a></td>
          <td><?php echo($fila['precio']); ?></td>
          <td><?php if($fila['unidades']==0){echo("Sin stock");}else{echo($fila['unidades']);} ?></td>
          <td><?php echo($fila['ventas']); ?></td>
          <td><?php echo($fila['prom_calificacion']); ?></td>

        </tr>

        <?php
        $fila = mysqli_fetch_array($res);
        }
        ?>


  </table>

  </div>

</body>
</html>
