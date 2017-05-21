/**
 * Created by Dapino on 21/05/17
 */

let map = null;
let tablaInstituciones = document.getElementById('tablaInstituciones');

function ajaxRequest() {
    let xhr = new XMLHttpRequest();

    xhr.open('GET', '/api/educativas', true)
    xhr.addEventListener('load', e => {
        let response = e.target,
            data = JSON.parse(response.responseText);

        document.querySelector('#preloader').classList.add('hide');
        initMap(data);
    });
    xhr.addEventListener('loadstart', () => {
        document.querySelector('#preloader').classList.remove('hide');
    });
    xhr.send();
}

function initMap(data) {
    map = new google.maps.Map(document.getElementById('map'), {
        zoom: 12,
        center: new google.maps.LatLng(4.5981259, -74.0782322),
        mapTypeId: 'roadmap'
    });
    let localidades = data.localidades;
    let instituciones = [];

    let labels = 'ABCDEFGHIJKLMNOPQRSTUVWXYZ';
    
    var markers = localidades.map((localidad, i) => {
        return new google.maps.Marker({
        position: {lat:localidad.latitud, lng: localidad.longitud},
        label: labels[i % labels.length]
        });
    });

    markers.forEach((marker, i) => {
        marker.addListener('click', () => {
            document.querySelector('#sidebarText').classList.add('hide');
            document.querySelector('#sidebarTable').classList.remove('hide');
            instituciones = data.result.records.filter(institucion => {
                return institucion.localidad === localidades[i].id;
            });
            tablaInstituciones.innerHTML = '';
          document.querySelector('#nombreLocalidad').textContent = localidades[i].nombre;
          instituciones.forEach(v => {
              let fila = document.createElement('tr');
              let cNombre = document.createElement('td');
              let cFormal = document.createElement('td');
              cNombre.textContent = v.nombreinstitucion;
              cFormal.textContent = v.formal;
              fila.appendChild(cNombre);
              fila.appendChild(cFormal);
              tablaInstituciones.appendChild(fila);
          })
        });
    })
    var markerCluster = new MarkerClusterer(map, markers,
        {imagePath: 'https://developers.google.com/maps/documentation/javascript/examples/markerclusterer/m'});
    

}
