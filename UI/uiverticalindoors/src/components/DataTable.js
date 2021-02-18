import * as React from 'react';
import { DataGrid } from '@material-ui/data-grid';
import  { useState,useEffect } from 'react';
import  DataServiceApi from '../services/DataServiceApi';



const DataTable =()=>{

  const [dataTable,setDataGrid] = useState([]);
const columns = [
  { field: 'id', headerName: 'ID Sensor', width: 130 },
  { field: 'temp', headerName: 'Temp ', width: 130 },
  { field: 'moisture', headerName: 'Moisture', width: 130 },
];



useEffect(() => {

   
  DataServiceApi.getListMetrics(setDataGrid);

},[]);

return (

  <div style={{ height: 400, width: '100%' }}>
    <DataGrid rows={dataTable} columns={columns} pageSize={5} checkboxSelection />
    <p>
    {console.log(dataTable)}
    </p>
  </div>
);

};
export default DataTable;
 

