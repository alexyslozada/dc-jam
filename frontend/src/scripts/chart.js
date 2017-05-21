/**
 * Created by alexys on 20/05/17.
 */
(() => {
    let data = {},
        declaradas = [],
        vigencias = [],
        expulsadas = [],
        recibidas = [];
    let ctx = document.getElementById("myChart");

    let xhr = new XMLHttpRequest();
    xhr.open('GET', '/api/desplazamiento', true)
    xhr.addEventListener('load', e => {
        let response = e.target;
        console.log(JSON.parse(response.responseText));
        data = JSON.parse(response.responseText);
        data.result.records.forEach((v, i) => {
            if (i > 0) {
                declaradas.push(v.declaradas);
                vigencias.push(v.VIGENCIA);
                expulsadas.push(v.expulsadas);
                recibidas.push(v.recibidas);
            }
        });
        console.log(vigencias);
        let myChart = new Chart(ctx, {
            type: 'bar',
            data: {
                labels: vigencias,
                datasets: [{
                        label: "Declaradas",
                        data: declaradas,
                        backgroundColor: 'red'
                    },
                    {
                        label: 'Expulsadas',
                        data: expulsadas,
                        backgroundColor: 'blue'
                    },
                    {
                        label: 'Recibidas',
                        data: recibidas,
                        backgroundColor: 'green'
                    }
                ]
            }
        });

    });
    xhr.send();

})();
