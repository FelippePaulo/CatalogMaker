
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
        postData(`http://localhost:8080/catalog`,{
            item: "item1",
            item2: "item2"
        })
    }
    
}

function loadCatalogs(catalog){
    console.log(catalog)
    divCatalog = document.createElement("div")
    divCatalog.classList.add("catalog")
    title = document.createElement("div")
    title.innerText = catalog.title
    content = document.createElement("div")
    content.innerText = catalog.description

    form = document.createElement("form")
    form.method = "GET"
    
    button = document.createElement("button") 
    button.innerText = "visitar"
    button.onclick = (e) => {
        e.preventDefault()
        loadProducts(catalog.id)
    }
    form.appendChild(button)

    divCatalog.appendChild(title)
    divCatalog.appendChild(content)
    divCatalog.appendChild(form)

    catalogContainer = document.getElementById("testeFetch")
    catalogContainer.appendChild(divCatalog)
}

function loadProducts(id){
    productContainer = document.getElementById("productContainer")
    products = document.createElement("div")
    productContainer.removeChild(productContainer.firstChild)
    productContainer.appendChild(products)
    getData(`http://localhost:8080/catalog/${id}`).then((result) => {
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

async function postData(url, obj){
    try {
        const response = await fetch(url, {
            method: 'POST',
            headers: { 
                'Content-Type': 'application/json' 
            },
            body: JSON.stringify(obj) // Certifique-se de incluir o corpo da solicitação
        });

        if (response.ok) {
            alert('Cadastrado com sucesso.');
        } else {
            throw new Error('Erro ' + response.status);
        }
    } catch (err) {
        console.error(err);
    }
}
