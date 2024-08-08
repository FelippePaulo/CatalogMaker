
function main(){
    catalogContainer = document.getElementById("testeFetch")
    catalogContainer.innerText = ""
    getData('http://localhost:8080/catalog').then((result) => {
        result.forEach(e => {
            loadCatalogs(e)
        });
        //catalogContainer.innerHTML = JSON.stringify(result)
        
    })

    //teste post
    submitButton = document.getElementById("postSubmit")
    submitButton.onclick = (event) => {
        event.preventDefault()
        postData(`http://localhost:8080/catalog`, {
            item1: document.getElementById("input1").value,
            item2: document.getElementById("input2").value,
            item3: document.getElementById("input3").value
        })
    }

    //teste put
    changeButton = document.getElementById("putSubmit")
    changeButton.onclick = (event) => {
        event.preventDefault()
        updateData(`http://localhost:8080/catalog`,document.getElementById("input0Put").value, {
            
            item1: document.getElementById("input1Put").value,
            item2: document.getElementById("input2Put").value,
            item3: document.getElementById("input3Put").value
        })
    }
    
}

function loadCatalogs(catalog){
    //console.log(catalog)
    divCatalog = document.createElement("div")
    divCatalog.classList.add("catalog")
    title = document.createElement("div")
    title.innerText = catalog.title
    content = document.createElement("div")
    content.innerText = catalog.description
    img = document.createElement("div")
    img.innerText = catalog.imgLink

    form = document.createElement("form")
    form.method = "GET"
    
    button = document.createElement("button") 
    button.innerText = "visitar"
    button.onclick = (e) => {
        e.preventDefault()
        loadProducts(catalog.id)
    }
    form.appendChild(button)

    deleteButton = document.createElement("button")
    deleteButton.innerText = "Excluir"
    deleteButton.onclick = (e) => {
        e.preventDefault()
        deleteData('http://localhost:8080/catalog',catalog.id)
    }
    form.appendChild(deleteButton)

    divCatalog.appendChild(title)
    divCatalog.appendChild(content)
    divCatalog.appendChild(img)
    divCatalog.appendChild(form)

    catalogContainer = document.getElementById("testeFetch")
    catalogContainer.appendChild(divCatalog)
}

function loadProducts(id){
    productContainer = document.getElementById("productContainer")
    products = document.createElement("div")
    productContainer.removeChild(productContainer.firstChild)
    productContainer.appendChild(products)
    getData(`http://localhost:8080/catalog/${id}/produto`).then((result) => {
        result.forEach(product => {
            console.log(product)
            divProduct = document.createElement("div")
            divProduct.classList.add("product")
            title = document.createElement("div")
            title.innerText = product.title

            divProduct.appendChild(title)

            products.appendChild(divProduct)
        });
    })
}

async function getData(url) {
    //const url = 'http://localhost:8080/catalog';
    try {
      const response = await fetch(url);

      if (!response.ok) {
        throw new Error(`Response status: ${response.status}`);
      }
  
      const json = await response.json();
      //console.log( json);
      return json
    } catch (error) {
      console.error(error.message);
    }
}

async function postData(url, obj) {
    //console.log(JSON.stringify(obj))
    try {
        const response = await fetch(url, {
            method: 'POST',
            headers: {
                'Content-Type': 'text/plain' 
            },
            body: JSON.stringify(obj) 
        });

        if (response.ok) {
            alert('Cadastrado com sucesso.');
        } else {
            const errorText = await response.text();
            throw new Error('Erro ' + response.status + ': ' + errorText);
        }
    } catch (err) {
        alert(err.message);
    }
}

async function updateData(url, id, obj){
    try {
        const response = await fetch(url + `/${id}`, {
            method: 'PUT',
            headers: {
                'Content-Type':  'application/json' 
            },
            body: JSON.stringify(obj) 
        });

        if (response.ok) {
            alert('Alterado com sucesso.');
        } else {
            const errorText = await response.text();
            throw new Error('Erro ' + response.status + ': ' + errorText);
        }
    } catch (err) {
        alert(err.message);
    }
}

async function deleteData(url, id){
    try {
        const response = await fetch(url + `/${id}`, {
            method: 'DELETE'
        });

        if (response.ok) {
            alert('Removido com sucesso.');
        } else {
            const errorText = await response.text();
            throw new Error('Erro ' + response.status + ': ' + errorText);
        }
    } catch (err) {
        alert(err.message);
    }
        
}
