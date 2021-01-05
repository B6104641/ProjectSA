import React, { useState, useEffect } from 'react';
import { makeStyles, Theme, createStyles } from '@material-ui/core/styles';
import {
  Content,
  Header,
  Page,
  pageTheme /*ContentHeader,*/,
} from '@backstage/core';
import { Link as RouterLink } from 'react-router-dom';
import SaveIcon from '@material-ui/icons/Save';
import AccountBoxIcon from '@material-ui/icons/AccountBox';
import { Grid, Container } from '@material-ui/core';
import Link from '@material-ui/core/Link';
import { DefaultApi } from '../../api/apis';
import TextField from '@material-ui/core/TextField';
import Button from '@material-ui/core/Button';
import FormControl from '@material-ui/core/FormControl';
import { Alert } from '@material-ui/lab';
import Select from '@material-ui/core/Select';
import InputLabel from '@material-ui/core/InputLabel';
import MenuItem from '@material-ui/core/MenuItem';
import InputAdornment from '@material-ui/core/InputAdornment';

import { EntUser } from '../../api/models/EntUser';
import { EntEquipment } from '../../api/models/EntEquipment';
import { EntCompany } from '../../api/models/EntCompany';

const useStyles = makeStyles((theme: Theme) =>
  createStyles({
    root: {
      display: 'flex',
      flexWrap: 'wrap',
      justifyContent: 'center',
    },
    margin: {
      margin: theme.spacing(3),
    },
    withoutLabel: {
      marginTop: theme.spacing(3),
    },
    paper: {
      marginTop: theme.spacing(2),
      marginBottom: theme.spacing(2),
    },
    button: {
      margin: theme.spacing(1),
    },
    textField: {
      width: '25ch',
    },
    selectEmpty: {
      marginTop: theme.spacing(2),
    },
    container: {
      display: 'flex',
      flexWrap: 'wrap',
    },
    formControl: {
      width: 300,
    },
    
  }),
);

/*let am: any;
let pp: any;
let sum: any = 0;
    equipments.map((item: any) => {
      (am = item.aMOUNT),
      (pp  = item.edges.equipment.eQUIPMENTPRICE),
      (sum = sum(am*pp))
    } */

export default function Create() {
  const classes = useStyles();
  const profile = { givenName: 'ระบบรถพยาบาล' };
  const api = new DefaultApi();
  //const [order, setOrder] = React.useState<Partial<inputOrder>>({});
  //const [user, setUser] = useState(initialUserState);
  const [equipments, setEquipments] = useState<EntEquipment[]>([]);
  const [companys, setCompanys] = useState<EntCompany[]>([]);
  const [users, setUsers] = useState<EntUser[]>([]);

  const [status, setStatus] = useState(false);
  const [alert, setAlert] = useState(true);
  const [loading, setLoading] = useState(true);

  const [equipmentid, setEquipmentid] = useState(Number);
  const [companyid, setCompanyid] = useState(Number);
  const [userid, setUserid] = useState(Number);
  const [amountg, setAmount] = useState(String);
  const [priceg, setPrice] = useState(Number);
  const [datetime, setDatetime] = useState(String);

  useEffect(() => {
    const getEquipments = async () => {
      const e = await api.listEquipment({ limit: 10, offset: 0 });
      setLoading(false);
      setEquipments(e);
    };
    getEquipments();

    const getCompanys = async () => {
      const cp = await api.listCompany({ limit: 10, offset: 0 });
      setLoading(false);
      setCompanys(cp);
    };
    getCompanys();

    const getUsers = async () => {
      const u = await api.listUser({ limit: 10, offset: 0 });
      setLoading(false);
      setUsers(u);
    };
    getUsers();
  }, [loading]);

/*let am : any;
let pp : any;
let sum : any = 0;

 /*equipments.map((item:any) => {

  (am = item.aMOUNT),
  (pp = item.edges.equipment.eQUIPMENTPRICE),
  (sum = sum+(am*pp) 
  )
}
); */

  const CreateOrder = async () => {
    const order = {
      equipment: equipmentid,
      company: companyid,
      user: userid,
      amount: amountg,
      price: priceg,
      added: datetime + ":00+07:00",
    };
    console.log(order);

    const res: any = await api.createOrder({ order: order });
    setStatus(true);
    if (res.id != '') {
      setAlert(true);
    } else {
      setAlert(false);
    }
   
  };

  const equipment_id_handleChange = (
    event: React.ChangeEvent<{ value: unknown }>,
  ) => {
    setEquipmentid(event.target.value as number);
  };

  const user_id_handleChange = (
    event: React.ChangeEvent<{ value: unknown }>,
  ) => {
    setUserid(event.target.value as number);
  };

  const company_id_handleChange = (
    event: React.ChangeEvent<{ value: unknown }>,
  ) => {
    setCompanyid(event.target.value as number);
  };

  const amoung_id_handleChange = (
    event: React.ChangeEvent<{ value: unknown }>,
  ) => {
    setAmount(event.target.value as string);
  };

  const priceg_id_handleChange = (
    event: React.ChangeEvent<{ value: unknown }>,
  ) => {
    setPrice(event.target.value as number);
  };
  
  const handleDatetimeChange = (event: any) => {
    setDatetime(event.target.value as string);
  };

  return (
    <Page theme={pageTheme.service}>
      <Header
        title={`${profile.givenName || 'to Backstage'}`}
        subtitle="ระบบสั่งซื้ออุปกรณ์."
      >
       
      </Header>
      <Content>
        <Container maxWidth="sm">
          <Grid container spacing={3}>
            <Grid item xs={12}></Grid>

            <Grid item xs={3}></Grid>
            <Grid item xs={9}>
              {status ? (
                <div>
                  {alert ? (
                    <Alert severity="success" style={{ width: 305 }}>
                      สั่งซื้อสำเร็จ :)
                    </Alert>
                  ) : (
                    <Alert severity="warning" style={{ marginTop: 20 }}>
                      ไม่สำเร็จ !
                    </Alert>
                  )}
                </div>
              ) : null}
            </Grid>

            <Grid item xs={3}>
              <div className={classes.paper}>อุปกรณ์</div>
            </Grid>
            <Grid item xs={9}>
              <FormControl variant="outlined" className={classes.formControl}>
                <InputLabel> </InputLabel>
                <Select
                  labelId="equipment_id-label"
                  //label="Equipment"
                  // style={{ width: 300 }}
                  id="equipment_id"
                  value={equipmentid}
                  onChange={equipment_id_handleChange}
                  
                >
                  {equipments.map((item: EntEquipment) => (
                    <MenuItem value={item.id}>{item.eQUIPMENTNAME}
                    </MenuItem>
                  ))}
                </Select>
              </FormControl>
            </Grid>

            <Grid item xs={3}>
              <div className={classes.paper}>บริษัท</div>
            </Grid>
            <Grid item xs={9}>
              <FormControl variant="outlined" className={classes.formControl}>
                <InputLabel> </InputLabel>
                <Select
                  labelId="company_id-label"
                  //label="Company"
                  //style={{ width: 300 }}
                  id="company_id"
                  value={companyid}
                  onChange={company_id_handleChange}
                >
                  {companys.map((item: EntCompany) => (
                    <MenuItem value={item.id}>{item.cOMPANYName}</MenuItem>
                  ))}
                </Select>
              </FormControl>
            </Grid>

            <Grid item xs={3}>
              <div className={classes.paper}>ผู้ดำเนินการ</div>
            </Grid>
            <Grid item xs={9}>
              <FormControl variant="outlined" className={classes.formControl}>
                <InputLabel> </InputLabel>
                <Select
                  labelId="user_id-label"
                  //label="User"
                  //style={{ width: 300 }}
                  id="user_id"
                  value={userid}
                  onChange={user_id_handleChange}
                >
                  {users.map((item: EntUser) => (
                    <MenuItem value={item.id}>{item.uSERNAME}</MenuItem>
                  ))}
                </Select>
              </FormControl>
            </Grid>

            <Grid item xs={3}>
              <div className={classes.paper}>จำนวน</div>
            </Grid>
            <Grid item xs={9}>
              <FormControl variant="outlined" className={classes.formControl}>
                <TextField
                  //style={{ width: 300 }}
                  id="outlined-number"
                  name="amount"
                  value={amountg}
                  type="number"
                  onChange={amoung_id_handleChange}
                  InputLabelProps={{
                    shrink: true,
                  }}
                  variant="outlined"
                />
              </FormControl>
            </Grid>

            <Grid item xs={3}>
              <div className={classes.paper}> ราคา/ชิ้น </div>
            </Grid>
            <Grid item xs={9}>
            <FormControl variant="outlined" className={classes.formControl}>
                <InputLabel> </InputLabel>
                <Select
                  labelId="equipment_id-label"
                  //label="Equipment"
                  // style={{ width: 300 }}
                  id="equipment_id"
                  value={priceg} 
                  onChange={priceg_id_handleChange}
                  startAdornment={<InputAdornment position="start">฿</InputAdornment>}
                >
                  {equipments.map((item: EntEquipment) => (
                    <MenuItem value={item.eQUIPMENTPRICE}>{item.eQUIPMENTPRICE}</MenuItem>
                  ))
                  }
                </Select>

              </FormControl>
            </Grid>

            <Grid item xs={3}>
              <div className={classes.paper}> แสดงราคา </div>
            </Grid>
            <Grid item xs={9}>
            <FormControl variant="outlined" className={classes.formControl}>
                <InputLabel> </InputLabel>
                <Select
                  labelId="equipment_id-label"
                  //label="Equipment"
                  // style={{ width: 300 }}
                  id="equipment_id"
                  value={equipmentid}
                  onChange={priceg_id_handleChange}
                  startAdornment={<InputAdornment position="start">฿</InputAdornment>}
                >
                  {equipments.map((item: EntEquipment) => (
                    <MenuItem value={item.id}>{item.eQUIPMENTPRICE}</MenuItem>
                  ))
                  }
                </Select>

              </FormControl>
            </Grid>

           

            <Grid item xs={3}>
              <div className={classes.paper}>วันที่สั่งซื้อ</div>
            </Grid>
            <Grid item xs={9}>
              <form className={classes.container} noValidate>
                <TextField
                  style={{ width: 300 }}
                  id="date"
                  //label="DateTime"
                  type="datetime-local"
                  value={datetime}
                  onChange={handleDatetimeChange}
                  //defaultValue="2017-05-24"
                  className={classes.textField}
                  InputLabelProps={{
                    shrink: true,
                  }}
                />
              </form>
            </Grid>

            <Grid item xs={3}></Grid>
            <Grid item xs={9}>
              <Button
                variant="contained"
                color="primary"
                size="small"
                className={classes.button}
                startIcon={<SaveIcon />}
                onClick={() => {
                  CreateOrder();
                }}
              >
                Save
              </Button>
              <Button
                size="small"
                style={{ marginLeft: 150 }}
                component={RouterLink}
                to="/result"
                variant="contained"
              >
                Back
              </Button>
            </Grid>
          </Grid>
        </Container>
      </Content>
    </Page>
  );
}
