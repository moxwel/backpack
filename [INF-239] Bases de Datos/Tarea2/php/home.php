<?php
session_start();
require "./database/db.php";

$query = "SELECT * FROM top5_prod_calif";
$res1 = mysqli_query($conn, $query);

$query = "SELECT * FROM top5_prod_venta";
$res2 = mysqli_query($conn, $query);

$query = "SELECT * FROM top5_ventas";
$res3 = mysqli_query($conn, $query);

$queryProd = "SELECT * FROM producto";
$resProd = mysqli_query($conn, $queryProd);

$queryCat = "SELECT * FROM etiqueta";
$resCat = mysqli_query($conn, $queryCat);
?>


<!DOCTYPE html>
<html>
<head>
<meta charset="UTF-8">
<meta name="viewport" content="width=device-width, initial-scale=1">
<title> SansApp </title>
<link rel="stylesheet" href="./extras/base.css">
<style>
  .column2 {
  float: left;
  width: 80%;
  padding: 15px;

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

<br>
<?php
  if (isset($_SESSION['mensaje'])) {
    echo("<center><strong>" . $_SESSION['mensaje'] . "</strong></center><br>");
    unset($_SESSION['mensaje']);
  }
?>
<h2>Pagina principal</h2>
<div class="row">
    <div class="column" style="background-color:lightblue;">
      <h2>Top 5 mejor calificados</h2>
      <?php

      $cont = 1;
      $fila1 = mysqli_fetch_array($res1);
      while (!is_null($fila1)) {
      ?>
      <p><?php echo($cont);?>.- <a href="producto.php?id=<?php echo($fila1['id_producto']); ?>"><?php echo($fila1['nombre']); ?></a> (<?php echo($fila1['prom_calificacion']); ?>)</p>

      <?php
      $cont = $cont + 1;
      $fila1 = mysqli_fetch_array($res1);
      }
      ?>

    </div>
    <div class="column" style="background-color:#E7A12F;">
      <h2>Top 5 más vendidos</h2>
      <?php

      $cont = 1;
      $fila2 = mysqli_fetch_array($res2);
      while (!is_null($fila2)) {
      ?>
      <p><?php echo($cont);?>.- <a href="producto.php?id=<?php echo($fila2['id_producto']); ?>"><?php echo($fila2['nombre']); ?></a> (<?php echo($fila2['ventas']); ?>)</p>

      <?php
      $cont = $cont + 1;
      $fila2 = mysqli_fetch_array($res2);
      }
      ?>

    </div>
    <div class="column" style="background-color:lightblue;">
      <h2>Top 5 usuarios con más ventas</h2>
      <?php

      $cont = 1;
      $fila3 = mysqli_fetch_array($res3);
      while (!is_null($fila3)) {
      ?>
      <p><?php echo($cont);?>.- <a href="perfil.php?id=<?php echo($fila3['id_cuenta']); ?>"><?php echo($fila3['nombre']); ?></a> (<?php echo($fila3['num_ventas']); ?>)</p>

      <?php
      $cont = $cont + 1;
      $fila3 = mysqli_fetch_array($res3);
      }
      ?>
    </div>
    <div class="column2" style="background-color:lightblue;width:75%;">
      <h2>Todos los productos</h2>
      <table>
        <tr>
          <th style="width:40%">Nombre producto</th>
          <th style="width:5%">Precio</th>
          <th style="width:5%">Cantidad</th>
          <th style="width:5%">Ventas</th>
          <th style="width:5%">Calificacion</th>

        </tr>


        <?php

        $filaProd = mysqli_fetch_array($resProd);
        while (!is_null($filaProd)) {
        ?>

        <tr>
          <td><a href="producto.php?id=<?php echo($filaProd['id_producto']); ?>"><?php echo($filaProd['nombre']); ?></a></td>
          <td><?php echo($filaProd['precio']); ?></td>
          <td><?php if($filaProd['unidades']==0){echo("Sin stock");}else{echo($filaProd['unidades']);} ?></td>
          <td><?php echo($filaProd['ventas']); ?></td>
          <td><?php echo($filaProd['prom_calificacion']); ?></td>

        </tr>

        <?php
        $filaProd = mysqli_fetch_array($resProd);
        }
        ?>


  </table>
  </div>
  <div class="column" style="background-color:#E7A12F;width:20%;">
      <h2>Categorias</h2>
      <?php

      $filaCat = mysqli_fetch_array($resCat);
      while (!is_null($filaCat)) {
      ?>
      <p>• <a href="buscarcategoria.php?q=<?php echo($filaCat['nombre']); ?>"><?php echo($filaCat['nombre']); ?></a></p>

      <?php
      $filaCat = mysqli_fetch_array($resCat);
      }
      ?>
    </div>
  </div>

</body>
</html>
