import axios from "axios";

export default async function getMP(tkn) {
    let config = {
        method: 'get',
        url: 'http://18.214.103.65:8080/api/plugins/telemetry/DEVICE/101d2fe0-454d-11ed-b4b1-1bcb8f5daa77/values/timeseries?keys=TIMESTAMP,WS,WD,Temp,RH,BP,Depth &startTs=1265046352083&endTs=1665043961492',
        headers: {
            'Content-Type': 'application/json',
            'X-Authorization': 'Bearer ' + tkn
        }
    };

    try {
        let resp = await axios(config);

        let fechas = resp.data.TIMESTAMP;
        let data_ws = resp.data.WS;
        let data_wd = resp.data.WD;
        let data_temp = resp.data.Temp;
        let data_rh = resp.data.RH;
        let data_bp = resp.data.BP;
        let datos = [ ];

        for(let i = 0; i < fechas.length; i++){
                    let arr = fechas[i].value.split(" ")

                    // console.log(data_temp[i].value)

                    datos[i] = {};

                    datos[i].ts = fechas[i].ts;
                    datos[i].date = arr[0];
                    datos[i].hour = arr[1];
                    datos[i].WS  = data_ws[i].value;
                    datos[i].WD  = data_wd[i].value;
                    datos[i].Temp = data_temp[i].value;
                    datos[i].RH  = data_rh[i].value;
                    datos[i].BP  = data_bp[i].value;
        }

        // console.log(datos);

        return datos;
    }
    catch (error) {
        return error.response.data;
    }
}
