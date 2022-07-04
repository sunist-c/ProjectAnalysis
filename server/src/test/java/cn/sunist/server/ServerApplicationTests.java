package cn.sunist.server;

import cn.sunist.server.bean.Country;
import cn.sunist.server.bean.Data;
import cn.sunist.server.bean.Province;
import cn.sunist.server.bean.Request;
import cn.sunist.server.dao.HiveJDBC;
import cn.sunist.server.util.DateFormat;
import com.alibaba.fastjson.JSON;
import org.junit.jupiter.api.Test;
import org.springframework.boot.test.context.SpringBootTest;
import org.springframework.kafka.annotation.EnableKafka;

import java.sql.SQLException;
import java.text.ParseException;
import java.util.ArrayList;
import java.util.List;

@SpringBootTest
class ServerApplicationTests {

    @Test
    void contextLoads() throws ParseException {
        Country china = new Country("China", "2020-01-26", "100", "13", "12");
        Province province = new Province("China", "Taiwan","2020-01-28", "24", "23", "22");
        System.out.println(JSON.toJSONString(china));
        System.out.println(JSON.toJSONString(province));
    }

    @Test
    void TestRequestCalculateCountry() throws ParseException, SQLException {
//        Request request = new Request("edwi", "2020-01-26", 2, 7);
//        System.out.printf(JSON.toJSONString(request));

        System.out.printf(JSON.toJSONString(HiveJDBC.calculateCountry(DateFormat.transform("2020-01-26"), 7)));
    }

    @Test
    void TestRequestCalculateProvince(){
        Request request = new Request("edwi", "2020-01-26", 3, 7);
        System.out.printf(JSON.toJSONString(request));
    }

    @Test
    void TestRequestAddData(){
        ArrayList<Data> objects = new ArrayList<>();
        objects.add(new Data("Chine","TaiWan","2020-01-26",1,10,24));
        objects.add(new Data("Chine","Nanjin","2020-01-26",2,3,4));
        Request request = new Request("enwfewf","2020-01-26",1,objects);
        System.out.printf(JSON.toJSONString(request));
    }

}
