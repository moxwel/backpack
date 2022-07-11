<div class="topnav">
  <a style="font-weight:bold;" class="active" href="/sansapp/php/home.php">SansApp</a>

  <c><label></label> </c>
  <form action="buscarproducto.php" method="GET">
    <d><input style="width:140px" type="text" placeholder="Buscar producto" name="q"></d>
    <a><button type="submit">🔎 Articulos</button></a>
  </form>
  <c><label></label> </c>
  <form action="buscarperfil.php" method="GET">
    <d><input style="width:140px" type="text" placeholder="Buscar perfil" name="q"></d>
    <a><button type="submit">🔎 Perfil</button></a>
  </form>
  <c><label></label> </c>
  <form action="buscarcategoria.php" method="GET">
    <d><input style="width:140px" type="text" placeholder="Buscar categoria" name="q"></d>
    <a><button type="submit">🔎 Categoria</button></a>
  </form>

  <?php

  if (isset($_SESSION['logged']) and $_SESSION['logged']) {
    $ses_id = $_SESSION['id'];
    $ses_nombre = $_SESSION['nombre'];
    $ses_rol = $_SESSION['rol_usm'];

    $queryNElementosCArro = "SELECT COUNT(*) AS n_carro FROM productos_carrito WHERE comprador='$ses_id'";
    $resNElementosCArro = mysqli_query($conn, $queryNElementosCArro);
    $filaNElementosCArro = mysqli_fetch_array($resNElementosCArro);

    $cant_carro = $filaNElementosCArro['n_carro'];
  ?>

<a style="float:right; font-weight: bold;" href="/sansapp/php/mi_perfil.php">👤 Perfil: <?php echo($ses_nombre); ?></a>
<a style="float:right; font-weight: bold;" href="/sansapp/php/carrito.php">🛒 Carrito (<?php echo($cant_carro); ?>)</a>

  <?php
  } else {
  ?>

  <a style="float:right; font-weight: bold;" href="/sansapp/php/login.php">🚪 Iniciar Sesion</a>

  <?php
  }
  ?>


</div>
