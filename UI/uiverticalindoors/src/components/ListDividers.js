import React from 'react';
import { makeStyles } from '@material-ui/core/styles';
import List from '@material-ui/core/List';
import ListItem from '@material-ui/core/ListItem';
import ListItemText from '@material-ui/core/ListItemText';
import Divider from '@material-ui/core/Divider';
import DataTable from './DataTable'
const useStyles = makeStyles((theme) => ({
  root: {
    width: '100%',
    maxWidth: 360,
    position: 'center',
    backgroundColor: theme.palette.background.paper,
  },
}));

export default function ListDividers() {
  const classes = useStyles();

const onClickAsesoria = ()=>{
    console.log("active link acesoria");
};

  return (
    <div> 
      <List component="nav" className={classes.root} aria-label="mailbox folders">
    <ListItem button>
      <ListItemText primary="Add new  Crops" />
    </ListItem>
    <Divider />
    <ListItem button divider>
      <ListItemText primary="Edit Racks" />
    </ListItem>
    <ListItem button>
      <ListItemText primary="Check Stats" 
      onClick ={onClickAsesoria}/>
    </ListItem>
    <Divider light />
    <ListItem button>
      <ListItemText primary="Remove Rack from Inventory" />
    </ListItem>
  </List>
  <DataTable></DataTable>
  </div>
  );
}