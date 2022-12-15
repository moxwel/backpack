import express from "express";
import getToken from "./backend/getToken.js";
import getRadiometro from "./backend/getRadiometro.js";
import getRadon from "./backend/getRadon.js";
import getMPGrimm from "./backend/getMPGrimm.js";
import getMP from "./backend/getMP.js";
import getBCMagee from "./backend/getBCMagee.js";
import getBCMAAP from "./backend/getBCMAAP.js";

const app = express();

var actualToken = ""

app.use("/authenticate", (req,res)=>{

  getToken().then((token) => {
    actualToken = token;
  })

  setTimeout(() => {
    res.redirect("/");

  }, 1000);

})

app.set("view engine", "ejs");

app.use(express.json());
app.use(express.static('public'))

app.get("/", (req, res) => {
  res.render("home", {
    routes: [
      { sensorName: "Radiómetro", link: "/viewdata/radio" },
      { sensorName: "Radón", link: "/viewdata/radon" },
      { sensorName: "MP-Grimm", link: "/viewdata/mpg" },
      { sensorName: "Multi-Parámetro", link: "/viewdata/mp" },
      { sensorName: "BC-Magee", link: "/viewdata/bcmagee" },
      { sensorName: "BC-MAAP", link: "/viewdata/bcmaap" },
    ],
    stringToken: actualToken
  });
});

app.get("/viewdata", (req, res) => {
  res.redirect("/")
})

app.get("/viewData/:sensor", (req, res) => {

  if (req.params.sensor === "radio") {
    getRadiometro(actualToken).then((response) => {
      res.render("tabla", { sensorName: "Radiómetro", data: response });
    });
  } else if (req.params.sensor === "radon") {
    getRadon(actualToken).then((response) => {
      res.render("tabla", { sensorName: "Radon", data: response });
    });
  } else if (req.params.sensor === "mpg") {
    getMPGrimm(actualToken).then((response) => {
      res.render("tablampg", { sensorName: "MP-Grimm", data: response });
    });
  } else if (req.params.sensor === "mp") {
    getMP(actualToken).then((response) => {
      res.render("tablamp", { sensorName: "Multi-Parametro", data: response });
    });
  } else if (req.params.sensor === "bcmagee") {
    getBCMagee(actualToken).then((response) => {
      res.render("tablabcmagee", { sensorName: "BC-Magee", data: response });
    });
  } else if (req.params.sensor === "bcmaap") {
    getBCMAAP(actualToken).then((response) => {
      res.render("tablabcmaap", { sensorName: "BC-MAAP", data: response });
    });
  } else {
    res.render("nopage");
  }
});

app.listen(3000, () => {
  console.log("Servidor hosteado en puerto 3000.\nhttp://localhost:3000");
});
