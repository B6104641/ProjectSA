  
import React, { useState, useEffect } from 'react';
import { makeStyles } from '@material-ui/core/styles';
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableContainer from '@material-ui/core/TableContainer';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';
import Paper from '@material-ui/core/Paper';
import Button from '@material-ui/core/Button';
import { DefaultApi } from '../../api/apis';
import { EntOrder } from '../../api/models/EntOrder';
import moment from 'moment';
 
const useStyles = makeStyles({
 table: {
   minWidth: 650,
 },
});
 
export default function ComponentsTable() {
 const classes = useStyles();
 const api = new DefaultApi();

 
 const [orders, setOrders] = useState<EntOrder[]>([]);
 const [loading, setLoading] = useState(true);
 
 useEffect(() => {
  const getPredicaments = async () => {
    const res = await api.listOrder({ limit: 10, offset: 0 });
    setLoading(false);
    setOrders(res);
    console.log(res);
  };
  getPredicaments();

}, [loading]);
 
 const deletePredicaments = async (id: number) => {
   const res = await api.deleteOrder({ id: id });
   setLoading(true);
 };
 
 return (
   <TableContainer component={Paper}>
     <Table className={classes.table} aria-label="simple table">
       <TableHead>
         <TableRow>
         <TableCell align="center">Order_No.</TableCell>
            <TableCell align="center">Equipment</TableCell>
            <TableCell align="center">Company</TableCell>
            <TableCell align="center">User</TableCell>
            <TableCell align="center">Amount</TableCell>
            <TableCell align="center">Price</TableCell>
            <TableCell align="center">Date</TableCell>
            <TableCell align="center">Manage</TableCell>
         </TableRow>
       </TableHead>
       <TableBody>
         {orders === undefined
            ? null
            : orders.map((item) => (
           <TableRow key={item.id}>

             <TableCell align="center">{item.id}</TableCell>
                  <TableCell align="center">{item.edges?.equipment?.eQUIPMENTNAME}</TableCell>
                  <TableCell align="center"> {item.edges?.company?.cOMPANYName} </TableCell>
                  <TableCell align="center">{item.edges?.user?.uSERNAME}</TableCell>
                  <TableCell align="center">{item.aMOUNT}</TableCell>
                  <TableCell align="center">{item.pRICE}</TableCell>
                  <TableCell align="center">{moment(item.aDDEDTIME).format('DD/MM/YYYY' + ' ' + 'HH:mm', )} </TableCell>
                  <TableCell align="center">
               <Button
                 onClick={() => {
                   if (item.id === undefined){
                     return;
                   }
                   deletePredicaments(item.id);
                 }}
                 style={{ marginLeft: 10 }}
                 variant="contained"
                 color="secondary"
               >
                 Delete
               </Button>
             </TableCell>
           </TableRow>
         ))}
       </TableBody>
     </Table>
   </TableContainer>
 );
}