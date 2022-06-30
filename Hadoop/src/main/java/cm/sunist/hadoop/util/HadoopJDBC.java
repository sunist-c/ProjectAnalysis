package cm.sunist.hadoop.util;

import cm.sunist.hadoop.bean.Data;
import org.apache.hadoop.conf.Configuration;
import org.apache.hadoop.fs.FSDataOutputStream;
import org.apache.hadoop.fs.FileSystem;
import org.apache.hadoop.fs.Path;

import java.io.IOException;
import java.net.URI;
import java.util.List;

public class HadoopJDBC {
    private static Configuration conf = null;
    private static String url = "hdfs://10.42.1.60:9000";
    private static FileSystem fs = null;
    static {
        Configuration configuration = new Configuration();
        configuration.set("fs.hdfs.impl", "org.apache.hadoop.hdfs.DistributedFileSystem");
        configuration.set("dfs.replication","3");
        try {
            fs = FileSystem.get(
                    URI.create(url),
                    configuration,
                    "root");
        } catch (IOException e) {
            e.printStackTrace();
        } catch (InterruptedException e) {
            e.printStackTrace();
        }
    }
    public static void createFile(String dst , List<Data> argList) throws IOException, InterruptedException, Exception{
        Path dstPath = new Path(url + dst); //目标路径
        //打开一个输出流
        FSDataOutputStream outputStream = fs.create(dstPath);
        StringBuffer sb = new StringBuffer();
        for(Data arg:argList){
            sb.append(arg.getLocationCountry()+","+arg.getLocationProvince()+","+arg.getRefreshTime()+","+arg.getConfirm()+","+arg.getDeath()+","+arg.getRecovered());
            sb.append("\n");
        }
        byte[] contents =  sb.toString().getBytes();
        outputStream.write(contents);
        outputStream.flush();;
        outputStream.close();
        System.out.println("文件创建成功！");
    }

    public static void deleteFile(String dst) throws IOException {
        Path dstPath = new Path(url + dst);
        fs.delete(dstPath,true);
    }
}
