# ship_ahoy
Alert when ships of interest are about to enter the part of the bay visible from our apartment.

Our apartment looks over part of the San Francisco bay. Alert if a ship of interest is about to enter that area. One side of that area is the part of the bay to the west of the Golden Gate bridge. The other part is the area to the east of Angel Island/Alcatraz.


# Implementation Notes and Background

Most recent ports of call.
https://www.vesselfinder.com/api/pro/portcalls/538007561?s

AIS vessel types.
https://help.marinetraffic.com/hc/en-us/articles/205579997-What-is-the-significance-of-the-AIS-Shiptype-number-

MMSI format.
https://www.navcen.uscg.gov/index.php?pageName=mtMmsi
https://en.wikipedia.org/wiki/Maritime_identification_digits
https://en.wikipedia.org/wiki/Maritime_Mobile_Service_Identity

MID registry.
https://www.itu.int/en/ITU-R/terrestrial/fmd/Pages/mid.aspx

Shine Micro DB.
http://www.mmsispace.com/livedisplay.php?mmsiresult=636091798
http://www.mmsispace.com/common/getdetails_v3.php?mmsi=369083000

Lat Lon calc.
https://www.movable-type.co.uk/scripts/latlong.html

# Data Requests

```
Ships in a region response record (from VesselFinder):
 22235849   -- lat * 600000
 522        -- lon * 600000
 683        -- Course *10
 117        -- Speed * 10
 280        -- ais
 249857000  -- mmsi
 WAIKIKI    -- Ship name
 0          -- unknown

MMSI data of a given ship
{
 'imo': '9776755',
 'name': 'WAIKIKI',
 'type': 'Crude Oil Tanker',
 't': '1531634521',
 'sar': False,
 'dest': 'MALTA',
 'etastamp': 'Jul 19, 12:00',
 'ship_speed': 11.7,
 'ship_course': 68.3,
 'timestamp': 'Jul 15, 2018 06:02 UTC',
 '__id': '309251',
 'pn': '9776755-249857000-6d17e98fedf50ed074675bd8f3396cd5',
 'vo': 0,
 'ff': False,
 'direct_link': '/vessels/WAIKIKI-IMO-9776755-MMSI-249857000',
 'draught': 8.8,
 'year': '2017',
 'gt': '61468',
 'sizes': '250 x 44 m',
 'dw': '112829'
}

Data query to ship_ahoy.ships
{
 'mmsi': '374518000',
 'imo': '8687218',
 'name': 'DONG HONG HANG 2',
 'ais': 170,
 'type': 'Bulk Carrier',
 'sar': 0,
 '__id': '0',
 'vo': 0,
 'ff': 0,
 'direct_link': '/vessels/DONG-HONG-HANG-2-IMO-8687218-MMSI-374518000',
 'draught': 5.0,
 'year': 2011,
 'gt': 8465,
 'sizes': '137 x 20 m',
 'length': 137,
 'beam': 20,
 'dw': 13685,
}
```

# SQL statements

```
CREATE TABLE ships (
    mmsi varchar(20),
    imo varchar(20),
    name varchar(128),
    ais int,
    type varchar(128),
    -- t unixtimestamp,
    sar boolean,
    -- dest varchar(255),
    -- etastamp 'Jun 21, 07:30',
    -- ship_speed float,
    -- ship_course float,
    -- timestamp 'Jun 27, 2018 17:48 UTC',
    __id varchar(20),
    -- pn varchar(255),  -- '0-227616590-808bd5b15abc2089364f4d3ccf1e13d6'
    vo int,
    ff boolean,
    direct_link varchar(128),
    draught float,
    year int,
    gt int,
    sizes varchar(50),
    length int not null,
    beam int not null,
    dw int,
    unknown int
 );

 CREATE UNIQUE INDEX mmsi ON ships ( mmsi );

 DELETE FROM ships;

 ALTER TABLE ships MODIFY mmsi varchar(20);
 ALTER TABLE ships ADD ais int AFTER name;

 UPDATE ships SET length = 0 WHERE length IS NULL;
 ALTER TABLE ships MODIFY length INT NOT NULL;

 ALTER TABLE ships DROP COLUMN unknown;

 CREATE TABLE sightings (
    mmsi varchar(20),
    ship_course float,
    timestamp int,  # Unix datetime
    lat float,
    lon float,
    my_lat float,
    my_lon float
 );
```

# Backup / Restore

```
mysqldump -u ships -p db_name t1 > dump.sql
mysql -u ships -p db_name < dump.sql
```

# Tidal Information

https://tidesandcurrents.noaa.gov/api/

Presidio tidal sensors https://tidesandcurrents.noaa.gov/stationhome.html?id=9414290

Bay Bridge Air Gap sensors https://tidesandcurrents.noaa.gov/map/index.html?id=9414304

Example queries

https://tidesandcurrents.noaa.gov/api/datagetter?date=latest&station=9414290&product=datums&datum=mllw&units=english&time_zone=lst_ldt&application=web_services&format=xml
```
<data>
<datum n="MHHW" v="11.817"/>
<datum n="MHW" v="11.208"/>
<datum n="DTL" v="8.897"/>
<datum n="MTL" v="9.160"/>
<datum n="MSL" v="9.097"/>
<datum n="MLW" v="7.113"/>
<datum n="MLLW" v="5.976"/>
<datum n="GT" v="5.841"/>
<datum n="MN" v="4.095"/>
<datum n="DHQ" v="0.609"/>
<datum n="DLQ" v="1.137"/>
<datum n="NAVD" v="5.917"/>
<datum n="LWI" v="2.781"/>
<datum n="HWI" v="24.721"/>
</data>
```

## Mean Lower Low Water for Presidio

https://tidesandcurrents.noaa.gov/api/datagetter?date=latest&station=9414290&product=water_level&datum=mllw&units=english&time_zone=lst_ldt&application=web_services&format=xml
```
<data>
<metadata id="9414290" name="San Francisco" lat="37.8063" lon="-122.4659"/>
<observations>
<wl t="2018-10-22 17:00" v="1.458" s="0.062" f="0,0,0,0" q="p"/>
</observations>
</data>
```

## Air gap for Bay Bridge D-E span

https://tidesandcurrents.noaa.gov/api/datagetter?date=latest&station=9414304&product=air_gap&datum=mllw&units=english&time_zone=lst_ldt&application=web_services&format=xml
```
<data>
<metadata id="9414304" name="San Francisco-Oakland Bay Bridge Air Gap" lat="37.8044" lon="-122.3728"/>
<observations>
<ag t="2018-10-24 16:48" v="204.400" s="0.121" f="1,0,0,0"/>
</observations>
</data>
```
