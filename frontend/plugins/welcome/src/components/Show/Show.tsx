import React, { FC } from 'react';
import { Link as RouterLink } from 'react-router-dom';
import ComponanceTable from '../Table';
import Button from '@material-ui/core/Button';
import AccountBoxIcon from '@material-ui/icons/AccountBox';
import Link from '@material-ui/core/Link';
import {Content, Header, Page, pageTheme, ContentHeader} from '@backstage/core';
 
const Show: FC<{}> = () => {
 const profile = { givenName: 'เข้าสู่ระบบรถพยาบาล' };
 
 return (
   <Page theme={pageTheme.service}>
     <Header
       title={`ยินดีต้อนรับ ${profile.givenName || 'to Backstage'}`}
       subtitle="ระบบย่อย: สั่งซื้ออุปกรณ์"  
     >
       <td align ='center'>
                  <AccountBoxIcon style={{ fontSize: 100}}>AccountBoxIcon</AccountBoxIcon>
                  <br></br>
                    gamse0505@gmail.com
                  <br></br>
                  <u>
                  <Link
                    component={RouterLink} to="/">
                        
                    Logout
                  </Link>
                  </u>
        </td>
     </Header>
     <Content>
       <ContentHeader title="ข้อมูลการใบสั่งซื้อ">
         <Link component={RouterLink} to="/user">
           <Button variant="contained" color="primary">
             สั่งซื้ออุปกรณ์
           </Button>
         </Link>
       </ContentHeader>
       <ComponanceTable></ComponanceTable>
     </Content>
   </Page>
 );
};
 
export default Show;