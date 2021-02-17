const DataServiceApi ={


getListMetrics : async function(setDataGrid){


    const url = "metricsgetlist";

    
   // const getresponse = async()=>{
    let headers = new Headers();
    headers.append("Access-Control-Allow-Origin",'*');
    headers.append('Access-Control-Allow-Methods', 'GET, POST, PUT','DELETE');
    
         await fetch(url,{
           headers
            
            }).then(response =>response.text())
              .then(rdata => {
                setDataGrid( rdata==="" ?[]: JSON.parse(rdata) )
              });
                
      //  };
    }
};
export default DataServiceApi;




