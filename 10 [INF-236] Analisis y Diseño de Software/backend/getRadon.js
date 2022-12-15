import axios from "axios";


export default async function getRadon(tkn) {
    let config = {
        method: 'get',
        url: 'http://18.214.103.65:8080/api/plugins/telemetry/DEVICE/723d0580-452d-11ed-b4b1-1bcb8f5daa77/values/timeseries?keys=Fecha,Radon&startTs=1265046352083&endTs=1665029708303',
        headers: {
            'Content-Type': 'application/json',
            'X-Authorization': 'Bearer ' + tkn
        }
    };

    try {
        let resp = await axios(config);

        let fechas = resp.data.Fecha;
        let datos  = resp.data.Radon;

        for(let i = 0; i < datos.length; i++){
            if(datos[i].value === '/'){
                datos[i].value = '-1';
            }
            let arr = fechas[i].value.split(" ")
            datos[i].date = arr[0];
            datos[i].hour = arr[1];
        }

        return datos;
    }
    catch (error) {
        return error.response.data;
    }
}
