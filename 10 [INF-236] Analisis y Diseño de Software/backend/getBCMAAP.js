import axios from "axios";

export default async function getBCMAAP(tkn) {
    let config = {
        method: 'get',
        url: 'http://18.214.103.65:8080/api/plugins/telemetry/DEVICE/6a4dd7a0-4550-11ed-b4b1-1bcb8f5daa77/values/timeseries?keys=fecha,hora,BC&startTs=1265046352083&endTs=1665044739122',
        headers: {
            'Content-Type': 'application/json',
            'X-Authorization': 'Bearer ' + tkn
        }
    };

    try {
        let resp = await axios(config);

        let fechas = resp.data.fecha;
        let horas = resp.data.hora;
        let data_bc = resp.data.BC;

        let datos = [ ];

        for(let i = 0; i < fechas.length; i++){

                    datos[i] = {};

                    datos[i].ts = fechas[i].ts;
                    datos[i].date = fechas[i].value;
                    datos[i].hour = horas[i].value;
                    datos[i].BC = data_bc[i].value;
        }

        // console.log(datos);

        return datos;
    }
    catch (error) {
        return error.response.data;
    }
}
