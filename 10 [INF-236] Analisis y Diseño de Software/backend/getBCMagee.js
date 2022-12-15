import axios from "axios";

export default async function getBCMagee(tkn) {
    let config = {
        method: 'get',
        url: 'http://18.214.103.65:8080/api/plugins/telemetry/DEVICE/8c5ad790-454f-11ed-b4b1-1bcb8f5daa77/values/timeseries?keys=Date,Time,BC1,BC2,BC3,BC4,BC5,BC6,BC7&startTs=1265046352083&endTs=1665044358966',
        headers: {
            'Content-Type': 'application/json',
            'X-Authorization': 'Bearer ' + tkn
        }
    };

    try {
        let resp = await axios(config);

        let fechas = resp.data.Date;
        let horas = resp.data.Time;
        let data_bc1 = resp.data.BC1;
        let data_bc2 = resp.data.BC2;
        let data_bc3 = resp.data.BC3;
        let data_bc4 = resp.data.BC4;
        let data_bc5 = resp.data.BC5;
        let data_bc6 = resp.data.BC6;
        let data_bc7 = resp.data.BC7;

        let datos = [ ];

        for(let i = 0; i < fechas.length; i++){

                    datos[i] = {};

                    datos[i].ts = fechas[i].ts;
                    datos[i].date = fechas[i].value;
                    datos[i].hour = horas[i].value;
                    datos[i].BC1 = data_bc1[i].value;
                    datos[i].BC2 = data_bc2[i].value;
                    datos[i].BC3 = data_bc3[i].value;
                    datos[i].BC4 = data_bc4[i].value;
                    datos[i].BC5 = data_bc5[i].value;
                    datos[i].BC6 = data_bc6[i].value;
                    datos[i].BC7 = data_bc7[i].value;
        }

        // console.log(datos);

        return datos;
    }
    catch (error) {
        return error.response.data;
    }
}
