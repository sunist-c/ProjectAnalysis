package cm.sunist.hadoop.bean;

import java.util.Objects;

public class Data {
    private String locationCountry;
    private String locationProvince;
    private String refreshTime;
    private Integer confirm;
    private Integer death;
    private Integer recovered;

    @Override
    public String toString() {
        return "Data{" +
                "locationCountry='" + locationCountry + '\'' +
                ", locationProvince='" + locationProvince + '\'' +
                ", refreshTime='" + refreshTime + '\'' +
                ", confirm=" + confirm +
                ", death=" + death +
                ", recovered=" + recovered +
                '}';
    }

    @Override
    public boolean equals(Object o) {
        if (this == o) return true;
        if (o == null || getClass() != o.getClass()) return false;
        Data data = (Data) o;
        return Objects.equals(locationCountry, data.locationCountry) && Objects.equals(locationProvince, data.locationProvince) && Objects.equals(refreshTime, data.refreshTime) && Objects.equals(confirm, data.confirm) && Objects.equals(death, data.death) && Objects.equals(recovered, data.recovered);
    }

    @Override
    public int hashCode() {
        return Objects.hash(locationCountry, locationProvince, refreshTime, confirm, death, recovered);
    }

    public String getLocationCountry() {
        return locationCountry;
    }

    public void setLocationCountry(String locationCountry) {
        this.locationCountry = locationCountry;
    }

    public String getLocationProvince() {
        return locationProvince;
    }

    public void setLocationProvince(String locationProvince) {
        this.locationProvince = locationProvince;
    }

    public String getRefreshTime() {
        return refreshTime;
    }

    public void setRefreshTime(String refreshTime) {
        this.refreshTime = refreshTime;
    }

    public Integer getConfirm() {
        return confirm;
    }

    public void setConfirm(Integer confirm) {
        this.confirm = confirm;
    }

    public Integer getDeath() {
        return death;
    }

    public void setDeath(Integer death) {
        this.death = death;
    }

    public Integer getRecovered() {
        return recovered;
    }

    public void setRecovered(Integer recovered) {
        this.recovered = recovered;
    }

    public Data() {
    }

    public Data(String locationCountry, String locationProvince, String refreshTime, Integer confirm, Integer death, Integer recovered) {
        this.locationCountry = locationCountry;
        this.locationProvince = locationProvince;
        this.refreshTime = refreshTime;
        this.confirm = confirm;
        this.death = death;
        this.recovered = recovered;
    }
}
