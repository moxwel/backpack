<?php
session_start();
require "./database/db.php";

if (!isset($_SESSION['logged'])) {
  unset($_SESSION['mensaje']);
  header('Location: /sansapp/php/home.php');
  die();
}
?>

<!DOCTYPE html>
<html>
<head>
<meta charset="UTF-8">
<meta name="viewport" content="width=device-width, initial-scale=1">
<title> Agregar producto </title>
<style>
Body {
  font-family: Calibri, Helvetica, sans-serif;
  background-color: #4BC3C7;
}
button {
       background-color: #E7A12F;
       width: 100%;
        color: white;
        padding: 15px;
        margin: 10px 0px;
        border: none;
        cursor: pointer;
         }
 form {
        border: 3px solid #f1f1f1;
    }
 input[type=text], input[type=number] {
        width: 100%;
        margin: 8px 0;
        padding: 12px 20px;
        display: inline-block;
        border: 2px solid lightblue;
        box-sizing: border-box;
    }
 button:hover {
        opacity: 0.7;
    }
  .opposition {
        width: auto;
        padding: 10px 18px;
        margin: 10px 5px;
    }


 .container {
        padding: 25px;
        background-color: lightblue;
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
    <form action="/sansapp/php/database/publish.php" method="post">
        <div class="container">
            <center> <h1> Publicar producto </h1> </center>
            <?php
                if (isset($_SESSION['mensaje'])) {
                echo("<center><strong>" . $_SESSION['mensaje'] . "</strong></center><br>");
                unset($_SESSION['mensaje']);
                }
            ?>
            <label>💬 Nombre del producto : </label>
            <input type="text" maxlength="20" placeholder="Introduce el nombre del producto" name="nombre" required>
            <label>💲 Precio : </label>
            <input type="number" min="0" max="9999999" placeholder="Introduce el precio del producto" name="precio" required>
            <label>🏷 Etiqueta : </label>
            <input type="text" maxlength="20" placeholder="Introduce la etiqueta del producto" name="etiqueta" required>
            <label>📦 Cantidad : </label>
            <input type="number" min="1" max="99999999999" placeholder="Introduce la cantidad que tienes del producto" name="cantidad" required>
            <label>📄 Descripcion : </label><br>
            <textarea name="descripcion" rows="20" cols="200" placeholder="Introduce una descripcion del producto" required></textarea>
            <!--<label>Contraseña : </label>
            <input type="contraseña" placeholder="Introduce tu contraseña" name="contraseña" required>   -->
            <button type="submit" name="env">✔ Agregar producto</button>

            <hr style="width:75%;">
            <center><button type="button" class="opposition"> <a style="color:white;" href="mis_productos.php">Atras</a></button> </center>
        </div>
    </form>
</body>
</html>
