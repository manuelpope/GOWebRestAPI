import { withStyles } from "@material-ui/core";
import { makeStyles } from '@material-ui/core/styles';

const useStyles = makeStyles((theme) => ({
    root: {
      width: '100%',
      maxWidth: 360,
      position: 'center',
      backgroundColor: theme.palette.background.paper,
    },
  }));
  
const HomePage =()=> {



    return(

        <h1>here it goes image , logo and home info</h1>
    );

}
export default (withStyles(useStyles) (HomePage));