package cm.sunist.hadoop.util;

import java.text.ParseException;
import java.text.SimpleDateFormat;
import java.util.Date;

public class DateFormat {
    private static SimpleDateFormat simpleDateFormat = new SimpleDateFormat("yyyy-mm-dd");
    public static String transform(Date date){
        return simpleDateFormat.format(date);
    }

    public static Date transform(String date) throws ParseException {
        return simpleDateFormat.parse(date);
    }
}
