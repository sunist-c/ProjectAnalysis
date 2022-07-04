package cn.sunist.server.bean;

import java.util.Objects;

public class Response {
    private Integer date;
    private String uuid;
    private Object data;

    public Response() {
    }

    public Response(Integer date, String uuid, Object data) {
        this.date = date;
        this.uuid = uuid;
        this.data = data;
    }

    @Override
    public String toString() {
        return "Response{" +
                "date=" + date +
                ", uuid='" + uuid + '\'' +
                ", data=" + data +
                '}';
    }

    @Override
    public boolean equals(Object o) {
        if (this == o) return true;
        if (o == null || getClass() != o.getClass()) return false;
        Response response = (Response) o;
        return Objects.equals(date, response.date) && Objects.equals(uuid, response.uuid) && Objects.equals(data, response.data);
    }

    @Override
    public int hashCode() {
        return Objects.hash(date, uuid, data);
    }

    public Integer getDate() {
        return date;
    }

    public void setDate(Integer date) {
        this.date = date;
    }

    public String getUuid() {
        return uuid;
    }

    public void setUuid(String uuid) {
        this.uuid = uuid;
    }

    public Object getData() {
        return data;
    }

    public void setData(Object data) {
        this.data = data;
    }
}
