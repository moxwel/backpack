import axios from "axios";

export default async function getMPGrimm(tkn) {
    let config = {
        method: 'get',
        url: 'http://18.214.103.65:8080/api/plugins/telemetry/DEVICE/99ce9f80-4557-11ed-b4b1-1bcb8f5daa77/values/timeseries?keys=Fecha,TSP,PM10,PM4,PM2.5,PM1&startTs=1265046352083&endTs=1665048457821',
        headers: {
            'Content-Type': 'application/json',
            'X-Authorization': 'Bearer ' + tkn
        }
    };

    try {
        let resp = await axios(config);

        let fechas = resp.data.Fecha;
        let data_tsp = resp.data.TSP;
        let data_pm10 = resp.data.PM10;
        let data_pm4 = resp.data.PM4;
        let data_pm25 = resp.data['PM2.5'];
        let data_pm1 = resp.data.PM1
        let datos = [ ];



        for(let i = 0; i < fechas.length; i++){
                    let arr = fechas[i].value.split(" ")

                    datos[i] = {};
                    datos[i].ts = fechas[i].ts;
                    datos[i].date = arr[0];
                    datos[i].hour = arr[1];
                    datos[i].TSP = data_tsp[i].value;
                    datos[i].PM10 = data_pm10[i].value;
                    datos[i].PM4 = data_pm4[i].value;
                    datos[i].PM25 = data_pm25[i].value;
                    datos[i].PM1 = data_pm1[i].value;

                }

        // console.log(datos);

        return datos;
    }
    catch (error) {
        return error.response.data;
    }
}
