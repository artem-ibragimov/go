package automaniac

const main_html = `
<!DOCTYPE html>
<html lang="en">
	<head>
		<meta charset="utf-8">
		<title>
			Specifications per brand and model - AutoManiac		</title>
		<meta content="width=device-width, initial-scale=1.0" name="viewport">
		<meta content="cars,prices,specs,information,advise,opinions,buying,selling,brands,models,failures,consumption,acceleration" name="keywords">
		<meta content="Detail information on cars and engines" name="description">
		<meta http-equiv="Cache-Control" content="no-cache, max-age=0, must-revalidate, no-store">
		<meta property="og:site_name" content="AutoManiac"/>
		<meta property="og:image" content="/resources/images/gui/cover_image.png" />
		<meta property="og:image:width" content="700"/>
		<meta property="og:image:height" content="300"/>
		<meta property="og:description" content="Detail information on cars and engines" />
		<meta property="og:title" content="Specifications per brand and model" />
		<meta property="fb:app_id" content="290604621342699" />
		<meta property="fb:admins" content="1276696339"/>
		<link rel="image_src" type="image/jpeg" href="/resources/images/gui/cover_image.png" />
		<link rel="shortcut icon" href="/resources/images/gui/favicon.png" type="image/x-icon">
		<link rel="icon" href="/resources/images/gui/favicon.png" type="image/x-icon">
			<script type="text/javascript" src="/resources/js/jquery-1.8.3.js"></script>
	<script type="text/javascript" src="/resources/js/jquery-ui.js"></script>
	<script type="text/javascript" src="/resources/js/scripts.js?version=v2.1.0"></script>
	<script type="text/javascript" src="/resources/js/jquery.chosen.min.js"></script>
	<script src="https://apis.google.com/js/platform.js" async defer></script>
	<script async src="https://securepubads.g.doubleclick.net/tag/js/gpt.js"></script>
	<!--
	<script async src="//pagead2.googlesyndication.com/pagead/js/adsbygoogle.js"></script>
	<script type="text/javascript" src="/resources/js/adsense.js"></script>
	-->
	<script async defer crossorigin="anonymous" src="https://connect.facebook.net/en_GB/sdk.js#xfbml=1&version=v4.0"></script>
	
						
	<link rel="stylesheet" href="/resources/css/styles.css?version=v2.1.6" type="text/css" media="screen" charset="utf-8" />
	<link rel="stylesheet" href="/resources/css/chosen.css" type="text/css" media="screen" charset="utf-8" />
	<link rel="stylesheet" href="/resources/css/font-awesome.min.css" >
	<link rel="stylesheet" href="https://fonts.googleapis.com/css?family=Open+Sans:300,300i,400,400i,700,700i|Raleway:300,400,500,700,800|Montserrat:300,400,700">
		<link rel="icon" href="/resources/images/gui/favicon.png" >
		<link rel="shortcut icon" href="/resources/images/gui/favicon.png" type="image/x-icon">
		<link rel="apple-touch-icon" href="images/gui/apple-touch-icon.png">
	</head>

	<body root-url="">

		<!-- Load Facebook SDK for JavaScript -->
		<div id="fb-root"></div>
		<script>
		  window.fbAsyncInit = function() {
		    FB.init({
		      appId      : '290604621342699',
		      xfbml      : true,
		      version    : 'v2.8'
		    });
		  };

		  (function(d, s, id){
		     var js, fjs = d.getElementsByTagName(s)[0];
		     if (d.getElementById(id)) {return;}
		     js = d.createElement(s); js.id = id;
		     js.src = "//connect.facebook.net/en_US/sdk.js";
		     fjs.parentNode.insertBefore(js, fjs);
		   }(document, 'script', 'facebook-jssdk'));

		  (function(w,d,u,h,s){
		    h=d.getElementsByTagName('head')[0];
		    s=d.createElement('script');
		    s.async=1;
		    s.src=u+'/sdk.js';
		    h.appendChild(s);
		  })(window,document,'https://aff.carvertical.com');

			
			  (function(i,s,o,g,r,a,m){i['GoogleAnalyticsObject']=r;i[r]=i[r]||function(){
			  (i[r].q=i[r].q||[]).push(arguments)},i[r].l=1*new Date();a=s.createElement(o),
			  m=s.getElementsByTagName(o)[0];a.async=1;a.src=g;m.parentNode.insertBefore(a,m)
			  })(window,document,'script','//www.google-analytics.com/analytics.js','ga');
			
			  ga('create', 'UA-75322623-2', 'auto');
	  		  ga('send', 'pageview');

	  		
		
		</script>

		<div class="menu-wrap">
			<div class="topnav" id="myTopnav">
				<div class="top-bar">
					<div class="top-bar-wrap">
						<div class="top-bar-left">
							<a href="/about" title="about">about</a>
							<a href="/legal" title="legal note">legal note</a>
							<a href="/contact" title="contact us">contact us</a>
							<span style="color: #fff; font-weight: bold;"></span>
						</div>
						<div class="top-bar-right">
							<a href="#" title="">
								Login <i class="fa fa-user" aria-hidden="true"></i>
							</a>
							<a href="https://www.automanie.org" title="">
								deutsch (D, AT, CH)<i class="fa fa-globe" aria-hidden="true"></i>
							</a>
							<a href="https://www.automanijak.com" title="">
								srpski (ex-yugoslavia)<i class="fa fa-flag" aria-hidden="true"></i>
							</a>
						</div>
					</div>
				</div>
				<a href="#" class="icon"><i class="fa fa-bars"></i></a>
				<div class="bottom-bar">
					<div class="bottom-bar-wrap">
						<div class="bottom-bar-left">
							<div class="logo">
								<a href="/" title="">
									<img src="/resources/images/gui/logo.png" alt="" title="">
								</a>
							</div>
						</div>
						<div class="bottom-bar-right">
							<a href="/virtual-adviser">Virtual adviser</a>
							<a href="/specs">Specifications by model</a>
							<a href="/compare">Compare two cars</a>
							<a href="/badges">Automotive badges</a>
							<a href="/blog">Blog</a>
						</div>
					</div>
				</div>
			</div>
		</div>
		<div id="main-wrapper">
			<div class="left-ad">
				<script async src="//pagead2.googlesyndication.com/pagead/js/adsbygoogle.js"></script>
<!-- branding -->
<ins class="adsbygoogle branding-ad"
     style="width:300px;height:600px"
     data-ad-client="ca-pub-2798966384669812"
     data-ad-slot="1676022217"></ins>
<script>
(adsbygoogle = window.adsbygoogle || []).push({});
</script>				<script async src="https://pagead2.googlesyndication.com/pagead/js/adsbygoogle.js"></script>
<!-- massive-header -->
<ins class="adsbygoogle massive-header-ad"
     data-ad-client="ca-pub-2798966384669812"
     data-ad-slot="9878901748"
     data-full-width-responsive="true"></ins>
<script>
     (adsbygoogle = window.adsbygoogle || []).push({});
</script>			</div>
			<div class="right-ad">
				<script async src="//pagead2.googlesyndication.com/pagead/js/adsbygoogle.js"></script>
<!-- branding -->
<ins class="adsbygoogle branding-ad"
     style="width:300px;height:600px"
     data-ad-client="ca-pub-2798966384669812"
     data-ad-slot="1676022217"></ins>
<script>
(adsbygoogle = window.adsbygoogle || []).push({});
</script>			</div>
			<div id="wrapper" class="clearfix">	<div id="sidebar">
<div id="sidebar-dodatno" class="search-box">
	<div class="sidebar-dodatno-wrap">
		<script async src="https://cse.google.com/cse.js?cx=016645478549947938819:d5ilzoxfr9e"></script>
		<script>
			window.onload = function(){
				$(".gsc-input").attr("placeholder", "Search this website...");
			}
		</script>
		<div class="gcse-search"></div>
	</div>
</div>	
	<div class="najpopulariji-wrap-bottom">
		<div class="sidebar-najpopularniji-naslov">
			<h2>Most popular models</h2>
		</div>
		<div class="najpop-item-wrap">
				
	<div class="najpop-item">
		<div class="najpop-item-name">
			<a href="/model/volkswagen/961/volkswagen-golf-1998" title="1998 Volkswagen Golf (Golf IV)">
				1998 Volkswagen Golf (Golf IV)			</a>
		</div>
		<div class="najpop-item-img">
			<a href="/model/volkswagen/961/volkswagen-golf-1998" title="1998 Volkswagen Golf (Golf IV)">
				<img src="/resources/images/variant/471/thumb_golf_1.jpg" 	title="Volkswagen Golf"
					 alt="Volkswagen Golf"
				 	 class="thumb-image">
			</a>
			<div class="ocena"><span>4.0</span>
			</div>
		</div>
		<div class="najpop-item-val">
			from: 			</br>
			<strong>900</strong>
			</br>EUR</div>
	</div>	
	<div class="najpop-item">
		<div class="najpop-item-name">
			<a href="/model/opel/694/opel-astra-1998" title="1998 Opel Astra (Astra G)">
				1998 Opel Astra (Astra G)			</a>
		</div>
		<div class="najpop-item-img">
			<a href="/model/opel/694/opel-astra-1998" title="1998 Opel Astra (Astra G)">
				<img src="/resources/images/variant/1337/thumb_astra_1.jpg" 	title="Opel Astra"
					 alt="Opel Astra"
				 	 class="thumb-image">
			</a>
			<div class="ocena"><span>4.1</span>
			</div>
		</div>
		<div class="najpop-item-val">
			from: 			</br>
			<strong>280</strong>
			</br>EUR</div>
	</div>	
	<div class="najpop-item">
		<div class="najpop-item-name">
			<a href="/model/bmw/255/bmw-3-series-2008" title="2008 BMW 3 Series (E90 rest...">
				2008 BMW 3 Series (E90 rest...			</a>
		</div>
		<div class="najpop-item-img">
			<a href="/model/bmw/255/bmw-3-series-2008" title="2008 BMW 3 Series (E90 rest...">
				<img src="/resources/images/variant/113/thumb_3_14.jpg" 	title="BMW 3 Series"
					 alt="BMW 3 Series"
				 	 class="thumb-image">
			</a>
			<div class="ocena"><span>4.1</span>
			</div>
		</div>
		<div class="najpop-item-val">
			from: 			</br>
			<strong>6.000</strong>
			</br>EUR</div>
	</div>	
	<div class="najpop-item">
		<div class="najpop-item-name">
			<a href="/model/mercedes-benz/630/mercedes-benz-e-2013" title="2013 Mercedes Benz E (W 212...">
				2013 Mercedes Benz E (W 212...			</a>
		</div>
		<div class="najpop-item-img">
			<a href="/model/mercedes-benz/630/mercedes-benz-e-2013" title="2013 Mercedes Benz E (W 212...">
				<img src="/resources/images/variant/1162/thumb_e_11.jpg" 	title="Mercedes Benz E"
					 alt="Mercedes Benz E"
				 	 class="thumb-image">
			</a>
			<div class="ocena"><span>4.0</span>
			</div>
		</div>
		<div class="najpop-item-val">
			from: 			</br>
			<strong>9.100</strong>
			</br>EUR</div>
	</div>	
	<div class="najpop-item">
		<div class="najpop-item-name">
			<a href="/model/škoda/1037/škoda-octavia-2009" title="2009 Škoda Octavia ">
				2009 Škoda Octavia 			</a>
		</div>
		<div class="najpop-item-img">
			<a href="/model/škoda/1037/škoda-octavia-2009" title="2009 Škoda Octavia ">
				<img src="/resources/images/variant/599/thumb_octavia_5.jpg" 	title="Škoda Octavia"
					 alt="Škoda Octavia"
				 	 class="thumb-image">
			</a>
			<div class="ocena"><span>4.6</span>
			</div>
		</div>
		<div class="najpop-item-val">
			from: 			</br>
			<strong>4.350</strong>
			</br>EUR</div>
	</div>	
	<div class="najpop-item">
		<div class="najpop-item-name">
			<a href="/model/audi/221/audi-a4-2004" title="2004 Audi A4 (B7)">
				2004 Audi A4 (B7)			</a>
		</div>
		<div class="najpop-item-img">
			<a href="/model/audi/221/audi-a4-2004" title="2004 Audi A4 (B7)">
				<img src="/resources/images/variant/41/thumb_a4_5.jpg" 	title="Audi A4"
					 alt="Audi A4"
				 	 class="thumb-image">
			</a>
			<div class="ocena"><span>4.0</span>
			</div>
		</div>
		<div class="najpop-item-val">
			from: 			</br>
			<strong>2.500</strong>
			</br>EUR</div>
	</div>	
	<div class="najpop-item">
		<div class="najpop-item-name">
			<a href="/model/renault/773/renault-megane-2002" title="2002 Renault Megane (Megane...">
				2002 Renault Megane (Megane...			</a>
		</div>
		<div class="najpop-item-img">
			<a href="/model/renault/773/renault-megane-2002" title="2002 Renault Megane (Megane...">
				<img src="/resources/images/variant/1429/thumb_megane_6.jpg" 	title="Renault Megane"
					 alt="Renault Megane"
				 	 class="thumb-image">
			</a>
			<div class="ocena"><span>4.3</span>
			</div>
		</div>
		<div class="najpop-item-val">
			from: 			</br>
			<strong>600</strong>
			</br>EUR</div>
	</div>	
	<div class="najpop-item">
		<div class="najpop-item-name">
			<a href="/model/peugeot/21/peugeot-206-2002" title="2002 Peugeot 206 ">
				2002 Peugeot 206 			</a>
		</div>
		<div class="najpop-item-img">
			<a href="/model/peugeot/21/peugeot-206-2002" title="2002 Peugeot 206 ">
				<img src="/resources/images/variant/395/thumb_206_3.jpg" 	title="Peugeot 206"
					 alt="Peugeot 206"
				 	 class="thumb-image">
			</a>
			<div class="ocena"><span>4.2</span>
			</div>
		</div>
		<div class="najpop-item-val">
			from: 			</br>
			<strong>650</strong>
			</br>EUR</div>
	</div>	
	<div class="najpop-item">
		<div class="najpop-item-name">
			<a href="/model/ford/388/ford-focus-2011" title="2011 Ford Focus ">
				2011 Ford Focus 			</a>
		</div>
		<div class="najpop-item-img">
			<a href="/model/ford/388/ford-focus-2011" title="2011 Ford Focus ">
				<img src="/resources/images/variant/756/thumb_focus_18.jpg" 	title="Ford Focus"
					 alt="Ford Focus"
				 	 class="thumb-image">
			</a>
			<div class="ocena"><span>4.5</span>
			</div>
		</div>
		<div class="najpop-item-val">
			from: 			</br>
			<strong>4.350</strong>
			</br>EUR</div>
	</div>	
	<div class="najpop-item">
		<div class="najpop-item-name">
			<a href="/model/fiat/56/fiat-500l-2012" title="2012 FIAT 500L ">
				2012 FIAT 500L 			</a>
		</div>
		<div class="najpop-item-img">
			<a href="/model/fiat/56/fiat-500l-2012" title="2012 FIAT 500L ">
				<img src="/resources/images/variant/633/thumb_500L_1.jpg" 	title="FIAT 500L"
					 alt="FIAT 500L"
				 	 class="thumb-image">
			</a>
			<div class="ocena"><span>4.2</span>
			</div>
		</div>
		<div class="najpop-item-val">
			from: 			</br>
			<strong>5.650</strong>
			</br>EUR</div>
	</div>		</div>
	</div>

	<div class="side-ad">
	<script async src="//pagead2.googlesyndication.com/pagead/js/adsbygoogle.js"></script>
	<ins class="adsbygoogle side-ad-vertical"
	     style="width:160px;height:600px"
	     data-ad-client="ca-pub-2798966384669812"
	     data-ad-slot="9565981886"></ins>
	<script>
	     (adsbygoogle = window.adsbygoogle || []).push({});
	</script>

	<ins class="adsbygoogle side-ad-horizontal"
	     data-ad-client="ca-pub-2798966384669812"
	     data-ad-slot="8347642282"
	     data-full-width-responsive="true"></ins>
	<script>
	     (adsbygoogle = window.adsbygoogle || []).push({});
	</script>
</div>
	</div>
	<div id="page-wrap">
<div id="breadcrumb-wrap">
	<div class="breadcrumb-nav"><a href='/'>HOME page</a> / specifications</div>
	<div id="sidebar-dodatno" class="search-box">
	<div class="sidebar-dodatno-wrap">
		<script async src="https://cse.google.com/cse.js?cx=016645478549947938819:d5ilzoxfr9e"></script>
		<script>
			window.onload = function(){
				$(".gsc-input").attr("placeholder", "Search this website...");
			}
		</script>
		<div class="gcse-search"></div>
	</div>
</div>	<div class="facebook">
				<div class="fb-like" data-href="https://www.facebook.com/automaniac.org" data-width="" data-layout="button_count" data-action="like" data-size="small" data-show-faces="true" data-share="true"></div>
	</div>
</div>	<div>
		<h2><i class="fa fa-car" aria-hidden="true"></i>Specifications per brand and model</h2>
	</div>
	<div id="modeli">
		<div class="box-wrap">
<div class='model'>		<a href='/make/alfa-romeo' title='Alfa Romeo'>			<img src='/resources/images/make/alfa-romeo.jpg' alt='Alfa Romeo Logo'/>		</a></div><div class='model'>		<a href='/make/audi' title='Audi'>			<img src='/resources/images/make/audi.jpg' alt='Audi Logo'/>		</a></div><div class='model'>		<a href='/make/bmw' title='BMW'>			<img src='/resources/images/make/bmw.jpg' alt='BMW Logo'/>		</a></div><div class='model'>		<a href='/make/chery' title='Chery'>			<img src='/resources/images/make/chery.jpg' alt='Chery Logo'/>		</a></div><div class='model'>		<a href='/make/chevrolet' title='Chevrolet'>			<img src='/resources/images/make/chevrolet.jpg' alt='Chevrolet Logo'/>		</a></div><div class='model'>		<a href='/make/chrysler' title='Chrysler'>			<img src='/resources/images/make/chrysler.jpg' alt='Chrysler Logo'/>		</a></div><div class='model'>		<a href='/make/citroen' title='Citroen'>			<img src='/resources/images/make/citroen.jpg' alt='Citroen Logo'/>		</a></div><div class='model'>		<a href='/make/dacia' title='Dacia'>			<img src='/resources/images/make/dacia.jpg' alt='Dacia Logo'/>		</a></div><div class='model'>		<a href='/make/daihatsu' title='Daihatsu'>			<img src='/resources/images/make/daihatsu.jpg' alt='Daihatsu Logo'/>		</a></div><div class='model'>		<a href='/make/dodge' title='Dodge'>			<img src='/resources/images/make/dodge.jpg' alt='Dodge Logo'/>		</a></div><div class='model'>		<a href='/make/fiat' title='FIAT'>			<img src='/resources/images/make/fiat.jpg' alt='FIAT Logo'/>		</a></div><div class='model'>		<a href='/make/ford' title='Ford'>			<img src='/resources/images/make/ford.jpg' alt='Ford Logo'/>		</a></div><div class='model'>		<a href='/make/honda' title='Honda'>			<img src='/resources/images/make/honda.jpg' alt='Honda Logo'/>		</a></div><div class='model'>		<a href='/make/hyundai' title='Hyundai'>			<img src='/resources/images/make/hyundai.jpg' alt='Hyundai Logo'/>		</a></div><div class='model'>		<a href='/make/infiniti' title='Infiniti'>			<img src='/resources/images/make/infiniti.jpg' alt='Infiniti Logo'/>		</a></div><div class='model'>		<a href='/make/jaguar' title='Jaguar'>			<img src='/resources/images/make/jaguar.jpg' alt='Jaguar Logo'/>		</a></div><div class='model'>		<a href='/make/jeep' title='Jeep'>			<img src='/resources/images/make/jeep.jpg' alt='Jeep Logo'/>		</a></div><div class='model'>		<a href='/make/kia' title='KIA'>			<img src='/resources/images/make/kia.jpg' alt='KIA Logo'/>		</a></div><div class='model'>		<a href='/make/lada' title='Lada'>			<img src='/resources/images/make/lada.jpg' alt='Lada Logo'/>		</a></div><div class='model'>		<a href='/make/lancia' title='Lancia'>			<img src='/resources/images/make/lancia.jpg' alt='Lancia Logo'/>		</a></div><div class='model'>		<a href='/make/land-rover' title='Land Rover'>			<img src='/resources/images/make/landrover.jpg' alt='Land Rover Logo'/>		</a></div><div class='model'>		<a href='/make/lexus' title='Lexus'>			<img src='/resources/images/make/lexus.jpg' alt='Lexus Logo'/>		</a></div><div class='model'>		<a href='/make/mazda' title='Mazda'>			<img src='/resources/images/make/mazda.jpg' alt='Mazda Logo'/>		</a></div><div class='model'>		<a href='/make/mercedes-benz' title='Mercedes Benz'>			<img src='/resources/images/make/mercedes.jpg' alt='Mercedes Benz Logo'/>		</a></div><div class='model'>		<a href='/make/mini' title='Mini'>			<img src='/resources/images/make/mini.jpg' alt='Mini Logo'/>		</a></div><div class='model'>		<a href='/make/mitsubishi' title='Mitsubishi'>			<img src='/resources/images/make/mitsubishi.jpg' alt='Mitsubishi Logo'/>		</a></div><div class='model'>		<a href='/make/nissan' title='Nissan'>			<img src='/resources/images/make/nissan.jpg' alt='Nissan Logo'/>		</a></div><div class='model'>		<a href='/make/opel' title='Opel'>			<img src='/resources/images/make/opel.jpg' alt='Opel Logo'/>		</a></div><div class='model'>		<a href='/make/peugeot' title='Peugeot'>			<img src='/resources/images/make/peugeot.jpg' alt='Peugeot Logo'/>		</a></div><div class='model'>		<a href='/make/porsche' title='Porsche'>			<img src='/resources/images/make/porsche.png' alt='Porsche Logo'/>		</a></div><div class='model'>		<a href='/make/renault' title='Renault'>			<img src='/resources/images/make/renault.jpg' alt='Renault Logo'/>		</a></div><div class='model'>		<a href='/make/rover' title='Rover'>			<img src='/resources/images/make/rover.jpg' alt='Rover Logo'/>		</a></div><div class='model'>		<a href='/make/saab' title='SAAB'>			<img src='/resources/images/make/saab.jpg' alt='SAAB Logo'/>		</a></div><div class='model'>		<a href='/make/ssangyong' title='SSangYong'>			<img src='/resources/images/make/ssang_yong.jpg' alt='SSangYong Logo'/>		</a></div><div class='model'>		<a href='/make/seat' title='Seat'>			<img src='/resources/images/make/seat.jpg' alt='Seat Logo'/>		</a></div><div class='model'>		<a href='/make/smart' title='Smart'>			<img src='/resources/images/make/smart.jpg' alt='Smart Logo'/>		</a></div><div class='model'>		<a href='/make/subaru' title='Subaru'>			<img src='/resources/images/make/subaru.jpg' alt='Subaru Logo'/>		</a></div><div class='model'>		<a href='/make/suzuki' title='Suzuki'>			<img src='/resources/images/make/suzuki.jpg' alt='Suzuki Logo'/>		</a></div><div class='model'>		<a href='/make/tesla' title='Tesla'>			<img src='/resources/images/make/tesla.jpg' alt='Tesla Logo'/>		</a></div><div class='model'>		<a href='/make/toyota' title='Toyota'>			<img src='/resources/images/make/toyota.png' alt='Toyota Logo'/>		</a></div><div class='model'>		<a href='/make/volkswagen' title='Volkswagen'>			<img src='/resources/images/make/vw.jpg' alt='Volkswagen Logo'/>		</a></div><div class='model'>		<a href='/make/volvo' title='Volvo'>			<img src='/resources/images/make/volvo.jpg' alt='Volvo Logo'/>		</a></div><div class='model'>		<a href='/make/zastava' title='Zastava'>			<img src='/resources/images/make/zastava.jpg' alt='Zastava Logo'/>		</a></div><div class='model'>		<a href='/make/škoda' title='Škoda'>			<img src='/resources/images/make/skoda.jpg' alt='Škoda Logo'/>		</a></div>		</div>
	</div><div>
	<h2><i class="fa fa-search" aria-hidden="true"></i>Check a car by its VIN number</h2>
</div>
<div class="box-wrap" style="height: 250px;">
	<div
	  data-cvaff
	  data-locale="en-EU"
	  data-a="automanijak"
	  data-b="81ec5429"
	  style="width: 100%; height: 100%">
	</div>
</div><div>
	<h2><i class="fa fa-globe" aria-hidden="true"></i>Follow us</h2>
</div>
<div id="facebook-statistika">
	<div class="box-wrap">
		<div class="box-left-wrap">
			<div class="fs-naslov">
				<h3>AutoManiac Instagram</h3>
			</div>
		</div>
		<div class="box-right-wrap"></div>
		<div id="curator-feed-default-feed-layout"><a href="https://curator.io" target="_blank" class="crt-logo crt-tag">Powered by Curator.io</a></div>
		<script type="text/javascript">
			(function(){
			var i,e,d=document,s="script";i=d.createElement("script");i.async=1;i.charset="UTF-8";
			i.src="https://cdn.curator.io/published/a532c6e2-2ba1-47aa-998d-99b1eb7101ab.js";
			e=d.getElementsByTagName(s)[0];e.parentNode.insertBefore(i, e);
			})();
		</script>
		<div class="box-left-wrap">
			<div class="fs-naslov">
				<h3>AutoManiac Facebook</h3>
			</div>
			<div class="fb-page" data-href="https://www.facebook.com/automaniac.org" data-tabs="" data-width="" data-height="" data-small-header="false" data-adapt-container-width="true" data-hide-cover="false" data-show-facepile="false">
				<blockquote cite="https://www.facebook.com/automaniac.org" class="fb-xfbml-parse-ignore">
					<a href="https://www.facebook.com/automaniac.org">AutoManiac</a>
				</blockquote>
			</div>
		</div>
		<div class="box-right-wrap">
			<div class="fs-naslov">
				<h3>AutoManiac database currently covers:</h3>
			</div>
			<div class="fs-stats"><span>44</span>worldwide automotive brands</div>
			<div class="fs-stats"><span>1.431</span>different vehicle models</div>
			<div class="fs-stats"><span>2.098</span>engines</div>
			<div class="fs-stats"><span>13.067</span>specific cars</div>
		</div>
	</div>
</div>	</div>
			</div>
		</div>
		<div id="footer">
			<div class="footer-wrap">
				Copyright &copy; 2015 - 2022 automaniac.org - All rights reserved.
				</br>
				Powered by <strong>Gama1 Solutions</strong>
				<div class="footer-links">
					<div class="footer-menu">
						<ul>
							<li><a href="/">HOME page</a></li>
							<li><a href="/about">about</a></li>
							<li><a href="/terms">terms & conditions</a></li>
							<li><a href="/legal">legal note</a></li>
							<li><a href="/faq">frequent questions (FAQ)</a></li>
							<li><a href="/contact">contact us</a></li>
						</ul>
					</div>

					<div class='footer-auto-brend'><ul><li><a href='/make/alfa-romeo'>Alfa Romeo</a></li><li><a href='/make/audi'>Audi</a></li><li><a href='/make/bmw'>BMW</a></li><li><a href='/make/chery'>Chery</a></li><li><a href='/make/chevrolet'>Chevrolet</a></li><li><a href='/make/chrysler'>Chrysler</a></li></ul></div><div class='footer-auto-brend'><ul><li><a href='/make/citroen'>Citroen</a></li><li><a href='/make/dacia'>Dacia</a></li><li><a href='/make/daihatsu'>Daihatsu</a></li><li><a href='/make/dodge'>Dodge</a></li><li><a href='/make/fiat'>FIAT</a></li><li><a href='/make/ford'>Ford</a></li></ul></div><div class='footer-auto-brend'><ul><li><a href='/make/honda'>Honda</a></li><li><a href='/make/hyundai'>Hyundai</a></li><li><a href='/make/infiniti'>Infiniti</a></li><li><a href='/make/jaguar'>Jaguar</a></li><li><a href='/make/jeep'>Jeep</a></li><li><a href='/make/kia'>KIA</a></li></ul></div><div class='footer-auto-brend'><ul><li><a href='/make/lada'>Lada</a></li><li><a href='/make/lancia'>Lancia</a></li><li><a href='/make/land-rover'>Land Rover</a></li><li><a href='/make/lexus'>Lexus</a></li><li><a href='/make/mazda'>Mazda</a></li><li><a href='/make/mercedes-benz'>Mercedes</a></li></ul></div><div class='footer-auto-brend'><ul><li><a href='/make/mini'>Mini</a></li><li><a href='/make/mitsubishi'>Mitsubishi</a></li><li><a href='/make/nissan'>Nissan</a></li><li><a href='/make/opel'>Opel</a></li><li><a href='/make/peugeot'>Peugeot</a></li><li><a href='/make/porsche'>Porsche</a></li></ul></div><div class='footer-auto-brend'><ul><li><a href='/make/renault'>Renault</a></li><li><a href='/make/rover'>Rover</a></li><li><a href='/make/saab'>SAAB</a></li><li><a href='/make/ssangyong'>SSangYong</a></li><li><a href='/make/seat'>Seat</a></li><li><a href='/make/smart'>Smart</a></li></ul></div><div class='footer-auto-brend'><ul><li><a href='/make/subaru'>Subaru</a></li><li><a href='/make/suzuki'>Suzuki</a></li><li><a href='/make/tesla'>Tesla</a></li><li><a href='/make/toyota'>Toyota</a></li><li><a href='/make/volkswagen'>Volkswagen</a></li><li><a href='/make/volvo'>Volvo</a></li></ul></div><div class='footer-auto-brend'><ul><li><a href='/make/zastava'>Zastava</a></li><li><a href='/make/škoda'>Škoda</a></li></ul></div>
				</div>
			</div>
		</div>
	</body>
</html>
`
const brand_html = `
<!DOCTYPE html>
<html lang="en">
	<head>
		<meta charset="utf-8">
		<title>
			Audi models - AutoManiac		</title>
		<meta content="width=device-width, initial-scale=1.0" name="viewport">
		<meta content="Audi specifications,Audi consumption,Audi acceleration,Audi prices,Audi buying,Audi selling,Audi opinions" name="keywords">
		<meta content="Audi Models - List of all Audi models and information about each one of them" name="description">
		<meta http-equiv="Cache-Control" content="no-cache, max-age=0, must-revalidate, no-store">
		<meta property="og:site_name" content="AutoManiac"/>
		<meta property="og:image" content="/resources/images/make/cover/audi.jpg" />
		<meta property="og:image:width" content="700"/>
		<meta property="og:image:height" content="300"/>
		<meta property="og:description" content="Audi Models - List of all Audi models and information about each one of them" />
		<meta property="og:title" content="Audi models" />
		<meta property="fb:app_id" content="290604621342699" />
		<meta property="fb:admins" content="1276696339"/>
		<link rel="image_src" type="image/jpeg" href="/resources/images/make/cover/audi.jpg" />
		<link rel="shortcut icon" href="/resources/images/gui/favicon.png" type="image/x-icon">
		<link rel="icon" href="/resources/images/gui/favicon.png" type="image/x-icon">
			<script type="text/javascript" src="/resources/js/jquery-1.8.3.js"></script>
	<script type="text/javascript" src="/resources/js/jquery-ui.js"></script>
	<script type="text/javascript" src="/resources/js/scripts.js?version=v2.1.0"></script>
	<script type="text/javascript" src="/resources/js/jquery.chosen.min.js"></script>
	<script src="https://apis.google.com/js/platform.js" async defer></script>
	<script async src="https://securepubads.g.doubleclick.net/tag/js/gpt.js"></script>
	<!--
	<script async src="//pagead2.googlesyndication.com/pagead/js/adsbygoogle.js"></script>
	<script type="text/javascript" src="/resources/js/adsense.js"></script>
	-->
	<script async defer crossorigin="anonymous" src="https://connect.facebook.net/en_GB/sdk.js#xfbml=1&version=v4.0"></script>
	
						
	<link rel="stylesheet" href="/resources/css/styles.css?version=v2.1.6" type="text/css" media="screen" charset="utf-8" />
	<link rel="stylesheet" href="/resources/css/chosen.css" type="text/css" media="screen" charset="utf-8" />
	<link rel="stylesheet" href="/resources/css/font-awesome.min.css" >
	<link rel="stylesheet" href="https://fonts.googleapis.com/css?family=Open+Sans:300,300i,400,400i,700,700i|Raleway:300,400,500,700,800|Montserrat:300,400,700">
		<link rel="icon" href="/resources/images/gui/favicon.png" >
		<link rel="shortcut icon" href="/resources/images/gui/favicon.png" type="image/x-icon">
		<link rel="apple-touch-icon" href="images/gui/apple-touch-icon.png">
	</head>

	<body root-url="">

		<!-- Load Facebook SDK for JavaScript -->
		<div id="fb-root"></div>
		<script>
		  window.fbAsyncInit = function() {
		    FB.init({
		      appId      : '290604621342699',
		      xfbml      : true,
		      version    : 'v2.8'
		    });
		  };

		  (function(d, s, id){
		     var js, fjs = d.getElementsByTagName(s)[0];
		     if (d.getElementById(id)) {return;}
		     js = d.createElement(s); js.id = id;
		     js.src = "//connect.facebook.net/en_US/sdk.js";
		     fjs.parentNode.insertBefore(js, fjs);
		   }(document, 'script', 'facebook-jssdk'));

		  (function(w,d,u,h,s){
		    h=d.getElementsByTagName('head')[0];
		    s=d.createElement('script');
		    s.async=1;
		    s.src=u+'/sdk.js';
		    h.appendChild(s);
		  })(window,document,'https://aff.carvertical.com');

			
			  (function(i,s,o,g,r,a,m){i['GoogleAnalyticsObject']=r;i[r]=i[r]||function(){
			  (i[r].q=i[r].q||[]).push(arguments)},i[r].l=1*new Date();a=s.createElement(o),
			  m=s.getElementsByTagName(o)[0];a.async=1;a.src=g;m.parentNode.insertBefore(a,m)
			  })(window,document,'script','//www.google-analytics.com/analytics.js','ga');
			
			  ga('create', 'UA-75322623-2', 'auto');
	  		  ga('send', 'pageview');

	  		
		
		</script>

		<div class="menu-wrap">
			<div class="topnav" id="myTopnav">
				<div class="top-bar">
					<div class="top-bar-wrap">
						<div class="top-bar-left">
							<a href="/about" title="about">about</a>
							<a href="/legal" title="legal note">legal note</a>
							<a href="/contact" title="contact us">contact us</a>
							<span style="color: #fff; font-weight: bold;"></span>
						</div>
						<div class="top-bar-right">
							<a href="#" title="">
								Login <i class="fa fa-user" aria-hidden="true"></i>
							</a>
							<a href="https://www.automanie.org" title="">
								deutsch (D, AT, CH)<i class="fa fa-globe" aria-hidden="true"></i>
							</a>
							<a href="https://www.automanijak.com" title="">
								srpski (ex-yugoslavia)<i class="fa fa-flag" aria-hidden="true"></i>
							</a>
						</div>
					</div>
				</div>
				<a href="#" class="icon"><i class="fa fa-bars"></i></a>
				<div class="bottom-bar">
					<div class="bottom-bar-wrap">
						<div class="bottom-bar-left">
							<div class="logo">
								<a href="/" title="">
									<img src="/resources/images/gui/logo.png" alt="" title="">
								</a>
							</div>
						</div>
						<div class="bottom-bar-right">
							<a href="/virtual-adviser">Virtual adviser</a>
							<a href="/specs">Specifications by model</a>
							<a href="/compare">Compare two cars</a>
							<a href="/badges">Automotive badges</a>
							<a href="/blog">Blog</a>
						</div>
					</div>
				</div>
			</div>
		</div>
		<div id="main-wrapper">
			<div class="left-ad">
				<script async src="//pagead2.googlesyndication.com/pagead/js/adsbygoogle.js"></script>
<!-- branding -->
<ins class="adsbygoogle branding-ad"
     style="width:300px;height:600px"
     data-ad-client="ca-pub-2798966384669812"
     data-ad-slot="1676022217"></ins>
<script>
(adsbygoogle = window.adsbygoogle || []).push({});
</script>				<script async src="https://pagead2.googlesyndication.com/pagead/js/adsbygoogle.js"></script>
<!-- massive-header -->
<ins class="adsbygoogle massive-header-ad"
     data-ad-client="ca-pub-2798966384669812"
     data-ad-slot="9878901748"
     data-full-width-responsive="true"></ins>
<script>
     (adsbygoogle = window.adsbygoogle || []).push({});
</script>			</div>
			<div class="right-ad">
				<script async src="//pagead2.googlesyndication.com/pagead/js/adsbygoogle.js"></script>
<!-- branding -->
<ins class="adsbygoogle branding-ad"
     style="width:300px;height:600px"
     data-ad-client="ca-pub-2798966384669812"
     data-ad-slot="1676022217"></ins>
<script>
(adsbygoogle = window.adsbygoogle || []).push({});
</script>			</div>
			<div id="wrapper" class="clearfix">	<div id="sidebar">
<div id="sidebar-dodatno" class="search-box">
	<div class="sidebar-dodatno-wrap">
		<script async src="https://cse.google.com/cse.js?cx=016645478549947938819:d5ilzoxfr9e"></script>
		<script>
			window.onload = function(){
				$(".gsc-input").attr("placeholder", "Search this website...");
			}
		</script>
		<div class="gcse-search"></div>
	</div>
</div>	
	<div class="najpopulariji-wrap-bottom">
		<div class="sidebar-najpopularniji-naslov">
			<h2>Most popular models</h2>
		</div>
		<div class="najpop-item-wrap">
				
	<div class="najpop-item">
		<div class="najpop-item-name">
			<a href="/model/audi/221/audi-a4-2004" title="2004 Audi A4 (B7)">
				2004 Audi A4 (B7)			</a>
		</div>
		<div class="najpop-item-img">
			<a href="/model/audi/221/audi-a4-2004" title="2004 Audi A4 (B7)">
				<img src="/resources/images/variant/41/thumb_a4_5.jpg" 	title="Audi A4"
					 alt="Audi A4"
				 	 class="thumb-image">
			</a>
			<div class="ocena"><span>4.0</span>
			</div>
		</div>
		<div class="najpop-item-val">
			from: 			</br>
			<strong>2.500</strong>
			</br>EUR</div>
	</div>	
	<div class="najpop-item">
		<div class="najpop-item-name">
			<a href="/model/audi/223/audi-a4-2011" title="2011 Audi A4 (B8 restyle)">
				2011 Audi A4 (B8 restyle)			</a>
		</div>
		<div class="najpop-item-img">
			<a href="/model/audi/223/audi-a4-2011" title="2011 Audi A4 (B8 restyle)">
				<img src="/resources/images/variant/46/thumb_a4_10.jpg" 	title="Audi A4"
					 alt="Audi A4"
				 	 class="thumb-image">
			</a>
			<div class="ocena"><span>4.3</span>
			</div>
		</div>
		<div class="najpop-item-val">
			from: 			</br>
			<strong>8.300</strong>
			</br>EUR</div>
	</div>	
	<div class="najpop-item">
		<div class="najpop-item-name">
			<a href="/model/audi/217/audi-a3-2008" title="2008 Audi A3 (8P restyle II)">
				2008 Audi A3 (8P restyle II)			</a>
		</div>
		<div class="najpop-item-img">
			<a href="/model/audi/217/audi-a3-2008" title="2008 Audi A3 (8P restyle II)">
				<img src="/resources/images/variant/31/thumb_a3_7.jpg" 	title="Audi A3"
					 alt="Audi A3"
				 	 class="thumb-image">
			</a>
			<div class="ocena"><span>4.0</span>
			</div>
		</div>
		<div class="najpop-item-val">
			from: 			</br>
			<strong>4.500</strong>
			</br>EUR</div>
	</div>	
	<div class="najpop-item">
		<div class="najpop-item-name">
			<a href="/model/audi/220/audi-a4-2001" title="2001 Audi A4 (B6)">
				2001 Audi A4 (B6)			</a>
		</div>
		<div class="najpop-item-img">
			<a href="/model/audi/220/audi-a4-2001" title="2001 Audi A4 (B6)">
				<img src="/resources/images/variant/39/thumb_a4_4.jpg" 	title="Audi A4"
					 alt="Audi A4"
				 	 class="thumb-image">
			</a>
			<div class="ocena"><span>4.0</span>
			</div>
		</div>
		<div class="najpop-item-val">
			from: 			</br>
			<strong>1.600</strong>
			</br>EUR</div>
	</div>	
	<div class="najpop-item">
		<div class="najpop-item-name">
			<a href="/model/audi/216/audi-a3-2005" title="2005 Audi A3 (8P restyle I)">
				2005 Audi A3 (8P restyle I)			</a>
		</div>
		<div class="najpop-item-img">
			<a href="/model/audi/216/audi-a3-2005" title="2005 Audi A3 (8P restyle I)">
				<img src="/resources/images/variant/29/thumb_a3_5.jpg" 	title="Audi A3"
					 alt="Audi A3"
				 	 class="thumb-image">
			</a>
			<div class="ocena"><span>4.4</span>
			</div>
		</div>
		<div class="najpop-item-val">
			from: 			</br>
			<strong>2.800</strong>
			</br>EUR</div>
	</div>		</div>
	</div>

	<div class="side-ad">
	<script async src="//pagead2.googlesyndication.com/pagead/js/adsbygoogle.js"></script>
	<ins class="adsbygoogle side-ad-vertical"
	     style="width:160px;height:600px"
	     data-ad-client="ca-pub-2798966384669812"
	     data-ad-slot="9565981886"></ins>
	<script>
	     (adsbygoogle = window.adsbygoogle || []).push({});
	</script>

	<ins class="adsbygoogle side-ad-horizontal"
	     data-ad-client="ca-pub-2798966384669812"
	     data-ad-slot="8347642282"
	     data-full-width-responsive="true"></ins>
	<script>
	     (adsbygoogle = window.adsbygoogle || []).push({});
	</script>
</div>
	</div>
	<div id="page-wrap">
<div id="breadcrumb-wrap">
	<div class="breadcrumb-nav"><a href='/'>HOME page</a> / <a href='/specs'>specifications</a> / Audi models</div>
	<div id="sidebar-dodatno" class="search-box">
	<div class="sidebar-dodatno-wrap">
		<script async src="https://cse.google.com/cse.js?cx=016645478549947938819:d5ilzoxfr9e"></script>
		<script>
			window.onload = function(){
				$(".gsc-input").attr("placeholder", "Search this website...");
			}
		</script>
		<div class="gcse-search"></div>
	</div>
</div>	<div class="facebook">
				<div class="fb-like" data-href="https://www.automaniac.org/make/audi" data-width="" data-layout="button_count" data-action="like" data-size="small" data-show-faces="true" data-share="true"></div>
	</div>
</div>	<div id="modeli-auto-lista">
		<div class="box-wrap">
			<div class="model-logo">
				<img src="/resources/images/make/audi.jpg"
					 alt="Audi models" />
			</div>
			<div class="model-naslov">
				Audi models				<span>List of all models from Audi</span>
			</div>
			<div class="model-zemlja">
				<img src="/resources/images/country/DE.png" alt="DE">
			</div>
			<div class="model-slika">
				<img src="/resources/images/make/cover/audi.jpg"
					 alt="Audi models" />
			</div>
		</div>
				<div>
			<h2><i class="fa fa-calendar" aria-hidden="true"></i>View by: <strong>segment</strong></h2>
		</div>
		<div class="box-wrap">
			<div class="segmenti">
				<div class="segment selected">
					<a href="/make/audi/segment"
					   title="View by<br/>segment"
					   alt="View by<br/>segment">
					<i class="fa fa-pie-chart" aria-hidden="true"></i>
						<p>View by<br/>segment</p>
					</a>
				</div>
				<div class="segment ">
					<a href="/make/audi/year"
					   title="View by<br/>year"
					   alt="View by<br/>year">
					<i class="fa fa-area-chart" aria-hidden="true"></i>
						<p>View by<br/>year</p>
					</a>
				</div>
				<div class="segment ">
					<a href="/make/audi/year-segment"
					   title="View by<br/>segment and year"
					   alt="View by<br/>segment and year">
						<i class="fa fa-bar-chart" aria-hidden="true"></i>
						<p>View by<br/>segment and year</p>
					</a>
				</div>
			</div>
		</div>


<div class="box-wrap">

	<div class="proizvodnja-wrap">

</div>
<div class="proizvodnja-wrap">


		<div class="proizvodnja-box variant-header accordion-controller segment-header"
			 related-accordion="accordion-B"
			 segment="B"
		 >
			<div class="proizvodnja-godina">
				B - City car			</div>
			<div class="proizvodnja-model">
				2 models			</div>
		</div>
		<div id="accordion-B" class="variant-accordion ">
							<div class="auto-wrap">
					<div class="auto-slika">
						<div class="auto-score-wrap clearfix">
														<div class="score-left">
								<div class="score4b">
									<p>drivers'</p>
									<span>3.6</span>
									<p>rating</p>
								</div>
							</div>
							<div class="score-right clearfix">
								<div class="score5a"></div>
								<div class="score4a"></div>
								<div class="score3a"></div>
								<div class="score2a"></div>
								<div class="score1a"></div>
							</div>
						</div>
						<a href="/model/audi/212/audi-a1-2010" title="">
							<img src="/resources/images/variant/22/thumb_a1_2.jpg"
								 alt=""
								 class="thumb-image"/>
						</a>
					</div>
					<div class="auto-info-wrap">
						<div class="auto-naziv">
							<a href="/model/audi/212/audi-a1-2010" title="2010 Audi A1 (8X)">
								2010 Audi A1 (8X)							</a>
						</div>
						<div class="stars">
							<div class="stars5"></div>
						</div>
						<div class="euro-ncap">
							EuroNCAP:&nbsp;
							90% occupant safety						</div>
						<div>
							<b>
								<a href="/model/audi/212/audi-a1-2010" title="2010 Audi A1 (8X)" class="orange">
							   	   	more details...						   	   </a>
					   	   </b>
						</div>
					</div>
				</div>
							<div class="auto-wrap">
					<div class="auto-slika">
						<div class="auto-score-wrap clearfix">
														<div class="score-left">
								<div class="score3b">
									<p>drivers'</p>
									<span>3.0</span>
									<p>rating</p>
								</div>
							</div>
							<div class="score-right clearfix">
								<div class="score5a"></div>
								<div class="score4a"></div>
								<div class="score3a"></div>
								<div class="score2a"></div>
								<div class="score1a"></div>
							</div>
						</div>
						<a href="/model/audi/1235/audi-a1-2014" title="">
							<img src="/resources/images/variant/1853/thumb_a1_3.jpg"
								 alt=""
								 class="thumb-image"/>
						</a>
					</div>
					<div class="auto-info-wrap">
						<div class="auto-naziv">
							<a href="/model/audi/1235/audi-a1-2014" title="2014 Audi A1 (8X restyle)">
								2014 Audi A1 (8X restyle)							</a>
						</div>
						<div class="stars">
							<div class="stars5"></div>
						</div>
						<div class="euro-ncap">
							EuroNCAP:&nbsp;
							90% occupant safety						</div>
						<div>
							<b>
								<a href="/model/audi/1235/audi-a1-2014" title="2014 Audi A1 (8X restyle)" class="orange">
							   	   	more details...						   	   </a>
					   	   </b>
						</div>
					</div>
				</div>
					</div>
</div>
<div class="proizvodnja-wrap">


		<div class="proizvodnja-box variant-header accordion-controller segment-header"
			 related-accordion="accordion-C"
			 segment="C"
		 >
			<div class="proizvodnja-godina">
				C - Small family car			</div>
			<div class="proizvodnja-model">
				7 models			</div>
		</div>
		<div id="accordion-C" class="variant-accordion ">
							<div class="auto-wrap">
					<div class="auto-slika">
						<div class="auto-score-wrap clearfix">
														<div class="score-left">
								<div class="score3b">
									<p>drivers'</p>
									<span>2.6</span>
									<p>rating</p>
								</div>
							</div>
							<div class="score-right clearfix">
								<div class="score5a"></div>
								<div class="score4a"></div>
								<div class="score3a"></div>
								<div class="score2a"></div>
								<div class="score1a"></div>
							</div>
						</div>
						<a href="/model/audi/214/audi-a3-1999" title="">
							<img src="/resources/images/variant/25/thumb_a3_1.jpg"
								 alt=""
								 class="thumb-image"/>
						</a>
					</div>
					<div class="auto-info-wrap">
						<div class="auto-naziv">
							<a href="/model/audi/214/audi-a3-1999" title="1999 Audi A3 (8L restyle)">
								1999 Audi A3 (8L restyle)							</a>
						</div>
						<div class="stars">
							<div class="stars4"></div>
						</div>
						<div class="euro-ncap">
							EuroNCAP:&nbsp;
							~70% occupant safety						</div>
						<div>
							<b>
								<a href="/model/audi/214/audi-a3-1999" title="1999 Audi A3 (8L restyle)" class="orange">
							   	   	more details...						   	   </a>
					   	   </b>
						</div>
					</div>
				</div>
							<div class="auto-wrap">
					<div class="auto-slika">
						<div class="auto-score-wrap clearfix">
														<div class="score-left">
								<div class="score4b">
									<p>drivers'</p>
									<span>4.1</span>
									<p>rating</p>
								</div>
							</div>
							<div class="score-right clearfix">
								<div class="score5a"></div>
								<div class="score4a"></div>
								<div class="score3a"></div>
								<div class="score2a"></div>
								<div class="score1a"></div>
							</div>
						</div>
						<a href="/model/audi/215/audi-a3-2003" title="">
							<img src="/resources/images/variant/27/thumb_a3_3.jpg"
								 alt=""
								 class="thumb-image"/>
						</a>
					</div>
					<div class="auto-info-wrap">
						<div class="auto-naziv">
							<a href="/model/audi/215/audi-a3-2003" title="2003 Audi A3 (8P)">
								2003 Audi A3 (8P)							</a>
						</div>
						<div class="stars">
							<div class="stars4"></div>
						</div>
						<div class="euro-ncap">
							EuroNCAP:&nbsp;
							~70% occupant safety						</div>
						<div>
							<b>
								<a href="/model/audi/215/audi-a3-2003" title="2003 Audi A3 (8P)" class="orange">
							   	   	more details...						   	   </a>
					   	   </b>
						</div>
					</div>
				</div>
							<div class="auto-wrap">
					<div class="auto-slika">
						<div class="auto-score-wrap clearfix">
														<div class="score-left">
								<div class="score4b">
									<p>drivers'</p>
									<span>4.4</span>
									<p>rating</p>
								</div>
							</div>
							<div class="score-right clearfix">
								<div class="score5a"></div>
								<div class="score4a"></div>
								<div class="score3a"></div>
								<div class="score2a"></div>
								<div class="score1a"></div>
							</div>
						</div>
						<a href="/model/audi/216/audi-a3-2005" title="">
							<img src="/resources/images/variant/29/thumb_a3_5.jpg"
								 alt=""
								 class="thumb-image"/>
						</a>
					</div>
					<div class="auto-info-wrap">
						<div class="auto-naziv">
							<a href="/model/audi/216/audi-a3-2005" title="2005 Audi A3 (8P restyle I)">
								2005 Audi A3 (8P restyle I)							</a>
						</div>
						<div class="stars">
							<div class="stars4"></div>
						</div>
						<div class="euro-ncap">
							EuroNCAP:&nbsp;
							~70% occupant safety						</div>
						<div>
							<b>
								<a href="/model/audi/216/audi-a3-2005" title="2005 Audi A3 (8P restyle I)" class="orange">
							   	   	more details...						   	   </a>
					   	   </b>
						</div>
					</div>
				</div>
							<div class="auto-wrap">
					<div class="auto-slika">
						<div class="auto-score-wrap clearfix">
														<div class="score-left">
								<div class="score4b">
									<p>drivers'</p>
									<span>4.0</span>
									<p>rating</p>
								</div>
							</div>
							<div class="score-right clearfix">
								<div class="score5a"></div>
								<div class="score4a"></div>
								<div class="score3a"></div>
								<div class="score2a"></div>
								<div class="score1a"></div>
							</div>
						</div>
						<a href="/model/audi/217/audi-a3-2008" title="">
							<img src="/resources/images/variant/31/thumb_a3_7.jpg"
								 alt=""
								 class="thumb-image"/>
						</a>
					</div>
					<div class="auto-info-wrap">
						<div class="auto-naziv">
							<a href="/model/audi/217/audi-a3-2008" title="2008 Audi A3 (8P restyle II)">
								2008 Audi A3 (8P restyle II)							</a>
						</div>
						<div class="stars">
							<div class="stars4"></div>
						</div>
						<div class="euro-ncap">
							EuroNCAP:&nbsp;
							~70% occupant safety						</div>
						<div>
							<b>
								<a href="/model/audi/217/audi-a3-2008" title="2008 Audi A3 (8P restyle II)" class="orange">
							   	   	more details...						   	   </a>
					   	   </b>
						</div>
					</div>
				</div>
							<div class="auto-wrap">
					<div class="auto-slika">
						<div class="auto-score-wrap clearfix">
														<div class="score-left">
								<div class="score5b">
									<p>drivers'</p>
									<span>4.6</span>
									<p>rating</p>
								</div>
							</div>
							<div class="score-right clearfix">
								<div class="score5a"></div>
								<div class="score4a"></div>
								<div class="score3a"></div>
								<div class="score2a"></div>
								<div class="score1a"></div>
							</div>
						</div>
						<a href="/model/audi/218/audi-a3-2012" title="">
							<img src="/resources/images/variant/34/thumb_a3_10.jpg"
								 alt=""
								 class="thumb-image"/>
						</a>
					</div>
					<div class="auto-info-wrap">
						<div class="auto-naziv">
							<a href="/model/audi/218/audi-a3-2012" title="2012 Audi A3 (8V)">
								2012 Audi A3 (8V)							</a>
						</div>
						<div class="stars">
							<div class="stars5"></div>
						</div>
						<div class="euro-ncap">
							EuroNCAP:&nbsp;
							95% occupant safety						</div>
						<div>
							<b>
								<a href="/model/audi/218/audi-a3-2012" title="2012 Audi A3 (8V)" class="orange">
							   	   	more details...						   	   </a>
					   	   </b>
						</div>
					</div>
				</div>
							<div class="auto-wrap">
					<div class="auto-slika">
						<div class="auto-score-wrap clearfix">
														<div class="score-left">
								<div class="score4b">
									<p>drivers'</p>
									<span>4.5</span>
									<p>rating</p>
								</div>
							</div>
							<div class="score-right clearfix">
								<div class="score5a"></div>
								<div class="score4a"></div>
								<div class="score3a"></div>
								<div class="score2a"></div>
								<div class="score1a"></div>
							</div>
						</div>
						<a href="/model/audi/1239/audi-a3-2016" title="">
							<img src="/resources/images/variant/1861/thumb_a3_13.jpg"
								 alt=""
								 class="thumb-image"/>
						</a>
					</div>
					<div class="auto-info-wrap">
						<div class="auto-naziv">
							<a href="/model/audi/1239/audi-a3-2016" title="2016 Audi A3 (8V restyle)">
								2016 Audi A3 (8V restyle)							</a>
						</div>
						<div class="stars">
							<div class="stars5"></div>
						</div>
						<div class="euro-ncap">
							EuroNCAP:&nbsp;
							95% occupant safety						</div>
						<div>
							<b>
								<a href="/model/audi/1239/audi-a3-2016" title="2016 Audi A3 (8V restyle)" class="orange">
							   	   	more details...						   	   </a>
					   	   </b>
						</div>
					</div>
				</div>
							<div class="auto-wrap">
					<div class="auto-slika">
						<div class="auto-score-wrap clearfix">
														<div class="score-left">
								<div class="score3b">
									<p>drivers'</p>
									<span>3.0</span>
									<p>rating</p>
								</div>
							</div>
							<div class="score-right clearfix">
								<div class="score5a"></div>
								<div class="score4a"></div>
								<div class="score3a"></div>
								<div class="score2a"></div>
								<div class="score1a"></div>
							</div>
						</div>
						<a href="/model/audi/1459/audi-a3-2020" title="">
							<img src="/resources/images/variant/2213/thumb_2020_audi_a3_sportback.jpg"
								 alt=""
								 class="thumb-image"/>
						</a>
					</div>
					<div class="auto-info-wrap">
						<div class="auto-naziv">
							<a href="/model/audi/1459/audi-a3-2020" title="2020 Audi A3 (8Y)">
								2020 Audi A3 (8Y)							</a>
						</div>
						<div class="stars">
							<div class="stars"></div>
						</div>
						<div class="euro-ncap">
							EuroNCAP:&nbsp;
							not tested						</div>
						<div>
							<b>
								<a href="/model/audi/1459/audi-a3-2020" title="2020 Audi A3 (8Y)" class="orange">
							   	   	more details...						   	   </a>
					   	   </b>
						</div>
					</div>
				</div>
					</div>
</div>
<div class="proizvodnja-wrap">


		<div class="proizvodnja-box variant-header accordion-controller segment-header"
			 related-accordion="accordion-D"
			 segment="D"
		 >
			<div class="proizvodnja-godina">
				D - Large family car			</div>
			<div class="proizvodnja-model">
				10 models			</div>
		</div>
		<div id="accordion-D" class="variant-accordion ">
							<div class="auto-wrap">
					<div class="auto-slika">
						<div class="auto-score-wrap clearfix">
														<div class="score-left">
								<div class="score4b">
									<p>drivers'</p>
									<span>4.2</span>
									<p>rating</p>
								</div>
							</div>
							<div class="score-right clearfix">
								<div class="score5a"></div>
								<div class="score4a"></div>
								<div class="score3a"></div>
								<div class="score2a"></div>
								<div class="score1a"></div>
							</div>
						</div>
						<a href="/model/audi/219/audi-a4-1999" title="">
							<img src="/resources/images/variant/37/thumb_a4_1.jpg"
								 alt=""
								 class="thumb-image"/>
						</a>
					</div>
					<div class="auto-info-wrap">
						<div class="auto-naziv">
							<a href="/model/audi/219/audi-a4-1999" title="1999 Audi A4 (B5 restyle)">
								1999 Audi A4 (B5 restyle)							</a>
						</div>
						<div class="stars">
							<div class="stars3"></div>
						</div>
						<div class="euro-ncap">
							EuroNCAP:&nbsp;
							~50% occupant safety						</div>
						<div>
							<b>
								<a href="/model/audi/219/audi-a4-1999" title="1999 Audi A4 (B5 restyle)" class="orange">
							   	   	more details...						   	   </a>
					   	   </b>
						</div>
					</div>
				</div>
							<div class="auto-wrap">
					<div class="auto-slika">
						<div class="auto-score-wrap clearfix">
														<div class="score-left">
								<div class="score4b">
									<p>drivers'</p>
									<span>4.0</span>
									<p>rating</p>
								</div>
							</div>
							<div class="score-right clearfix">
								<div class="score5a"></div>
								<div class="score4a"></div>
								<div class="score3a"></div>
								<div class="score2a"></div>
								<div class="score1a"></div>
							</div>
						</div>
						<a href="/model/audi/220/audi-a4-2001" title="">
							<img src="/resources/images/variant/39/thumb_a4_4.jpg"
								 alt=""
								 class="thumb-image"/>
						</a>
					</div>
					<div class="auto-info-wrap">
						<div class="auto-naziv">
							<a href="/model/audi/220/audi-a4-2001" title="2001 Audi A4 (B6)">
								2001 Audi A4 (B6)							</a>
						</div>
						<div class="stars">
							<div class="stars4"></div>
						</div>
						<div class="euro-ncap">
							EuroNCAP:&nbsp;
							~70% occupant safety						</div>
						<div>
							<b>
								<a href="/model/audi/220/audi-a4-2001" title="2001 Audi A4 (B6)" class="orange">
							   	   	more details...						   	   </a>
					   	   </b>
						</div>
					</div>
				</div>
							<div class="auto-wrap">
					<div class="auto-slika">
						<div class="auto-score-wrap clearfix">
														<div class="score-left">
								<div class="score4b">
									<p>drivers'</p>
									<span>4.0</span>
									<p>rating</p>
								</div>
							</div>
							<div class="score-right clearfix">
								<div class="score5a"></div>
								<div class="score4a"></div>
								<div class="score3a"></div>
								<div class="score2a"></div>
								<div class="score1a"></div>
							</div>
						</div>
						<a href="/model/audi/221/audi-a4-2004" title="">
							<img src="/resources/images/variant/41/thumb_a4_5.jpg"
								 alt=""
								 class="thumb-image"/>
						</a>
					</div>
					<div class="auto-info-wrap">
						<div class="auto-naziv">
							<a href="/model/audi/221/audi-a4-2004" title="2004 Audi A4 (B7)">
								2004 Audi A4 (B7)							</a>
						</div>
						<div class="stars">
							<div class="stars4"></div>
						</div>
						<div class="euro-ncap">
							EuroNCAP:&nbsp;
							~70% occupant safety						</div>
						<div>
							<b>
								<a href="/model/audi/221/audi-a4-2004" title="2004 Audi A4 (B7)" class="orange">
							   	   	more details...						   	   </a>
					   	   </b>
						</div>
					</div>
				</div>
							<div class="auto-wrap">
					<div class="auto-slika">
						<div class="auto-score-wrap clearfix">
														<div class="score-left">
								<div class="score4b">
									<p>drivers'</p>
									<span>4.0</span>
									<p>rating</p>
								</div>
							</div>
							<div class="score-right clearfix">
								<div class="score5a"></div>
								<div class="score4a"></div>
								<div class="score3a"></div>
								<div class="score2a"></div>
								<div class="score1a"></div>
							</div>
						</div>
						<a href="/model/audi/222/audi-a4-2007" title="">
							<img src="/resources/images/variant/44/thumb_a4_8.jpg"
								 alt=""
								 class="thumb-image"/>
						</a>
					</div>
					<div class="auto-info-wrap">
						<div class="auto-naziv">
							<a href="/model/audi/222/audi-a4-2007" title="2007 Audi A4 (B8)">
								2007 Audi A4 (B8)							</a>
						</div>
						<div class="stars">
							<div class="stars5"></div>
						</div>
						<div class="euro-ncap">
							EuroNCAP:&nbsp;
							~90% occupant safety						</div>
						<div>
							<b>
								<a href="/model/audi/222/audi-a4-2007" title="2007 Audi A4 (B8)" class="orange">
							   	   	more details...						   	   </a>
					   	   </b>
						</div>
					</div>
				</div>
							<div class="auto-wrap">
					<div class="auto-slika">
						<div class="auto-score-wrap clearfix">
														<div class="score-left">
								<div class="score4b">
									<p>drivers'</p>
									<span>4.3</span>
									<p>rating</p>
								</div>
							</div>
							<div class="score-right clearfix">
								<div class="score5a"></div>
								<div class="score4a"></div>
								<div class="score3a"></div>
								<div class="score2a"></div>
								<div class="score1a"></div>
							</div>
						</div>
						<a href="/model/audi/223/audi-a4-2011" title="">
							<img src="/resources/images/variant/46/thumb_a4_10.jpg"
								 alt=""
								 class="thumb-image"/>
						</a>
					</div>
					<div class="auto-info-wrap">
						<div class="auto-naziv">
							<a href="/model/audi/223/audi-a4-2011" title="2011 Audi A4 (B8 restyle)">
								2011 Audi A4 (B8 restyle)							</a>
						</div>
						<div class="stars">
							<div class="stars5"></div>
						</div>
						<div class="euro-ncap">
							EuroNCAP:&nbsp;
							93% occupant safety						</div>
						<div>
							<b>
								<a href="/model/audi/223/audi-a4-2011" title="2011 Audi A4 (B8 restyle)" class="orange">
							   	   	more details...						   	   </a>
					   	   </b>
						</div>
					</div>
				</div>
							<div class="auto-wrap">
					<div class="auto-slika">
						<div class="auto-score-wrap clearfix">
														<div class="score-left">
								<div class="score3b">
									<p>drivers'</p>
									<span>2.8</span>
									<p>rating</p>
								</div>
							</div>
							<div class="score-right clearfix">
								<div class="score5a"></div>
								<div class="score4a"></div>
								<div class="score3a"></div>
								<div class="score2a"></div>
								<div class="score1a"></div>
							</div>
						</div>
						<a href="/model/audi/1187/audi-a4-2015" title="">
							<img src="/resources/images/variant/1780/thumb_a4_25.jpg"
								 alt=""
								 class="thumb-image"/>
						</a>
					</div>
					<div class="auto-info-wrap">
						<div class="auto-naziv">
							<a href="/model/audi/1187/audi-a4-2015" title="2015 Audi A4 (B9)">
								2015 Audi A4 (B9)							</a>
						</div>
						<div class="stars">
							<div class="stars5"></div>
						</div>
						<div class="euro-ncap">
							EuroNCAP:&nbsp;
							90% occupant safety						</div>
						<div>
							<b>
								<a href="/model/audi/1187/audi-a4-2015" title="2015 Audi A4 (B9)" class="orange">
							   	   	more details...						   	   </a>
					   	   </b>
						</div>
					</div>
				</div>
							<div class="auto-wrap">
					<div class="auto-slika">
						<div class="auto-score-wrap clearfix">
														<div class="score-left">
								<div class="score3b">
									<p>drivers'</p>
									<span>3.0</span>
									<p>rating</p>
								</div>
							</div>
							<div class="score-right clearfix">
								<div class="score5a"></div>
								<div class="score4a"></div>
								<div class="score3a"></div>
								<div class="score2a"></div>
								<div class="score1a"></div>
							</div>
						</div>
						<a href="/model/audi/1394/audi-a4-2018" title="">
							<img src="/resources/images/variant/2107/thumb_2018_audi_a4.jpg"
								 alt=""
								 class="thumb-image"/>
						</a>
					</div>
					<div class="auto-info-wrap">
						<div class="auto-naziv">
							<a href="/model/audi/1394/audi-a4-2018" title="2018 Audi A4 (B9 restyle)">
								2018 Audi A4 (B9 restyle)							</a>
						</div>
						<div class="stars">
							<div class="stars5"></div>
						</div>
						<div class="euro-ncap">
							EuroNCAP:&nbsp;
							90% occupant safety						</div>
						<div>
							<b>
								<a href="/model/audi/1394/audi-a4-2018" title="2018 Audi A4 (B9 restyle)" class="orange">
							   	   	more details...						   	   </a>
					   	   </b>
						</div>
					</div>
				</div>
							<div class="auto-wrap">
					<div class="auto-slika">
						<div class="auto-score-wrap clearfix">
														<div class="score-left">
								<div class="score1b">
									<p>drivers'</p>
									<span>1.0</span>
									<p>rating</p>
								</div>
							</div>
							<div class="score-right clearfix">
								<div class="score5a"></div>
								<div class="score4a"></div>
								<div class="score3a"></div>
								<div class="score2a"></div>
								<div class="score1a"></div>
							</div>
						</div>
						<a href="/model/audi/226/audi-a5-sportback-2009" title="">
							<img src="/resources/images/variant/53/thumb_a5_5.jpg"
								 alt=""
								 class="thumb-image"/>
						</a>
					</div>
					<div class="auto-info-wrap">
						<div class="auto-naziv">
							<a href="/model/audi/226/audi-a5-sportback-2009" title="2009 Audi A5 Sportback (8T)">
								2009 Audi A5 Sportback (8T)							</a>
						</div>
						<div class="stars">
							<div class="stars"></div>
						</div>
						<div class="euro-ncap">
							EuroNCAP:&nbsp;
							not tested						</div>
						<div>
							<b>
								<a href="/model/audi/226/audi-a5-sportback-2009" title="2009 Audi A5 Sportback (8T)" class="orange">
							   	   	more details...						   	   </a>
					   	   </b>
						</div>
					</div>
				</div>
							<div class="auto-wrap">
					<div class="auto-slika">
						<div class="auto-score-wrap clearfix">
														<div class="score-left">
								<div class="score4b">
									<p>drivers'</p>
									<span>4.1</span>
									<p>rating</p>
								</div>
							</div>
							<div class="score-right clearfix">
								<div class="score5a"></div>
								<div class="score4a"></div>
								<div class="score3a"></div>
								<div class="score2a"></div>
								<div class="score1a"></div>
							</div>
						</div>
						<a href="/model/audi/227/audi-a5-sportback-2011" title="">
							<img src="/resources/images/variant/54/thumb_a5_6.jpg"
								 alt=""
								 class="thumb-image"/>
						</a>
					</div>
					<div class="auto-info-wrap">
						<div class="auto-naziv">
							<a href="/model/audi/227/audi-a5-sportback-2011" title="2011 Audi A5 Sportback (8T restyle)">
								2011 Audi A5 Sportback (8T restyle)							</a>
						</div>
						<div class="stars">
							<div class="stars"></div>
						</div>
						<div class="euro-ncap">
							EuroNCAP:&nbsp;
							not tested						</div>
						<div>
							<b>
								<a href="/model/audi/227/audi-a5-sportback-2011" title="2011 Audi A5 Sportback (8T restyle)" class="orange">
							   	   	more details...						   	   </a>
					   	   </b>
						</div>
					</div>
				</div>
							<div class="auto-wrap">
					<div class="auto-slika">
						<div class="auto-score-wrap clearfix">
														<div class="score-left">
								<div class="score3b">
									<p>drivers'</p>
									<span>3.0</span>
									<p>rating</p>
								</div>
							</div>
							<div class="score-right clearfix">
								<div class="score5a"></div>
								<div class="score4a"></div>
								<div class="score3a"></div>
								<div class="score2a"></div>
								<div class="score1a"></div>
							</div>
						</div>
						<a href="/model/audi/1245/audi-a5-sportback-2016" title="">
							<img src="/resources/images/variant/1877/thumb_a5_7.jpg"
								 alt=""
								 class="thumb-image"/>
						</a>
					</div>
					<div class="auto-info-wrap">
						<div class="auto-naziv">
							<a href="/model/audi/1245/audi-a5-sportback-2016" title="2016 Audi A5 Sportback (F5)">
								2016 Audi A5 Sportback (F5)							</a>
						</div>
						<div class="stars">
							<div class="stars5"></div>
						</div>
						<div class="euro-ncap">
							EuroNCAP:&nbsp;
							89% occupant safety						</div>
						<div>
							<b>
								<a href="/model/audi/1245/audi-a5-sportback-2016" title="2016 Audi A5 Sportback (F5)" class="orange">
							   	   	more details...						   	   </a>
					   	   </b>
						</div>
					</div>
				</div>
					</div>
</div>
<div class="proizvodnja-wrap">


		<div class="proizvodnja-box variant-header accordion-controller segment-header"
			 related-accordion="accordion-E"
			 segment="E"
		 >
			<div class="proizvodnja-godina">
				E - Luxury car			</div>
			<div class="proizvodnja-model">
				10 models			</div>
		</div>
		<div id="accordion-E" class="variant-accordion ">
							<div class="auto-wrap">
					<div class="auto-slika">
						<div class="auto-score-wrap clearfix">
														<div class="score-left">
								<div class="score4b">
									<p>drivers'</p>
									<span>4.1</span>
									<p>rating</p>
								</div>
							</div>
							<div class="score-right clearfix">
								<div class="score5a"></div>
								<div class="score4a"></div>
								<div class="score3a"></div>
								<div class="score2a"></div>
								<div class="score1a"></div>
							</div>
						</div>
						<a href="/model/audi/228/audi-a6-1997" title="">
							<img src="/resources/images/variant/55/thumb_a6_1.jpg"
								 alt=""
								 class="thumb-image"/>
						</a>
					</div>
					<div class="auto-info-wrap">
						<div class="auto-naziv">
							<a href="/model/audi/228/audi-a6-1997" title="1997 Audi A6 (C5)">
								1997 Audi A6 (C5)							</a>
						</div>
						<div class="stars">
							<div class="stars3"></div>
						</div>
						<div class="euro-ncap">
							EuroNCAP:&nbsp;
							~50% occupant safety						</div>
						<div>
							<b>
								<a href="/model/audi/228/audi-a6-1997" title="1997 Audi A6 (C5)" class="orange">
							   	   	more details...						   	   </a>
					   	   </b>
						</div>
					</div>
				</div>
							<div class="auto-wrap">
					<div class="auto-slika">
						<div class="auto-score-wrap clearfix">
														<div class="score-left">
								<div class="score4b">
									<p>drivers'</p>
									<span>4.4</span>
									<p>rating</p>
								</div>
							</div>
							<div class="score-right clearfix">
								<div class="score5a"></div>
								<div class="score4a"></div>
								<div class="score3a"></div>
								<div class="score2a"></div>
								<div class="score1a"></div>
							</div>
						</div>
						<a href="/model/audi/229/audi-a6-2001" title="">
							<img src="/resources/images/variant/57/thumb_a6_3.jpg"
								 alt=""
								 class="thumb-image"/>
						</a>
					</div>
					<div class="auto-info-wrap">
						<div class="auto-naziv">
							<a href="/model/audi/229/audi-a6-2001" title="2001 Audi A6 (C5 restyle)">
								2001 Audi A6 (C5 restyle)							</a>
						</div>
						<div class="stars">
							<div class="stars3"></div>
						</div>
						<div class="euro-ncap">
							EuroNCAP:&nbsp;
							~50% occupant safety						</div>
						<div>
							<b>
								<a href="/model/audi/229/audi-a6-2001" title="2001 Audi A6 (C5 restyle)" class="orange">
							   	   	more details...						   	   </a>
					   	   </b>
						</div>
					</div>
				</div>
							<div class="auto-wrap">
					<div class="auto-slika">
						<div class="auto-score-wrap clearfix">
														<div class="score-left">
								<div class="score4b">
									<p>drivers'</p>
									<span>4.2</span>
									<p>rating</p>
								</div>
							</div>
							<div class="score-right clearfix">
								<div class="score5a"></div>
								<div class="score4a"></div>
								<div class="score3a"></div>
								<div class="score2a"></div>
								<div class="score1a"></div>
							</div>
						</div>
						<a href="/model/audi/230/audi-a6-2004" title="">
							<img src="/resources/images/variant/59/thumb_a6_5.jpg"
								 alt=""
								 class="thumb-image"/>
						</a>
					</div>
					<div class="auto-info-wrap">
						<div class="auto-naziv">
							<a href="/model/audi/230/audi-a6-2004" title="2004 Audi A6 (C6)">
								2004 Audi A6 (C6)							</a>
						</div>
						<div class="stars">
							<div class="stars5"></div>
						</div>
						<div class="euro-ncap">
							EuroNCAP:&nbsp;
							~90% occupant safety						</div>
						<div>
							<b>
								<a href="/model/audi/230/audi-a6-2004" title="2004 Audi A6 (C6)" class="orange">
							   	   	more details...						   	   </a>
					   	   </b>
						</div>
					</div>
				</div>
							<div class="auto-wrap">
					<div class="auto-slika">
						<div class="auto-score-wrap clearfix">
														<div class="score-left">
								<div class="score4b">
									<p>drivers'</p>
									<span>4.4</span>
									<p>rating</p>
								</div>
							</div>
							<div class="score-right clearfix">
								<div class="score5a"></div>
								<div class="score4a"></div>
								<div class="score3a"></div>
								<div class="score2a"></div>
								<div class="score1a"></div>
							</div>
						</div>
						<a href="/model/audi/231/audi-a6-2008" title="">
							<img src="/resources/images/variant/61/thumb_a6_7.jpg"
								 alt=""
								 class="thumb-image"/>
						</a>
					</div>
					<div class="auto-info-wrap">
						<div class="auto-naziv">
							<a href="/model/audi/231/audi-a6-2008" title="2008 Audi A6 (C6 restyle)">
								2008 Audi A6 (C6 restyle)							</a>
						</div>
						<div class="stars">
							<div class="stars5"></div>
						</div>
						<div class="euro-ncap">
							EuroNCAP:&nbsp;
							~90% occupant safety						</div>
						<div>
							<b>
								<a href="/model/audi/231/audi-a6-2008" title="2008 Audi A6 (C6 restyle)" class="orange">
							   	   	more details...						   	   </a>
					   	   </b>
						</div>
					</div>
				</div>
							<div class="auto-wrap">
					<div class="auto-slika">
						<div class="auto-score-wrap clearfix">
														<div class="score-left">
								<div class="score4b">
									<p>drivers'</p>
									<span>4.1</span>
									<p>rating</p>
								</div>
							</div>
							<div class="score-right clearfix">
								<div class="score5a"></div>
								<div class="score4a"></div>
								<div class="score3a"></div>
								<div class="score2a"></div>
								<div class="score1a"></div>
							</div>
						</div>
						<a href="/model/audi/232/audi-a6-2011" title="">
							<img src="/resources/images/variant/63/thumb_a6_9.jpg"
								 alt=""
								 class="thumb-image"/>
						</a>
					</div>
					<div class="auto-info-wrap">
						<div class="auto-naziv">
							<a href="/model/audi/232/audi-a6-2011" title="2011 Audi A6 (C7)">
								2011 Audi A6 (C7)							</a>
						</div>
						<div class="stars">
							<div class="stars5"></div>
						</div>
						<div class="euro-ncap">
							EuroNCAP:&nbsp;
							91% occupant safety						</div>
						<div>
							<b>
								<a href="/model/audi/232/audi-a6-2011" title="2011 Audi A6 (C7)" class="orange">
							   	   	more details...						   	   </a>
					   	   </b>
						</div>
					</div>
				</div>
							<div class="auto-wrap">
					<div class="auto-slika">
						<div class="auto-score-wrap clearfix">
														<div class="score-left">
								<div class="score5b">
									<p>drivers'</p>
									<span>4.8</span>
									<p>rating</p>
								</div>
							</div>
							<div class="score-right clearfix">
								<div class="score5a"></div>
								<div class="score4a"></div>
								<div class="score3a"></div>
								<div class="score2a"></div>
								<div class="score1a"></div>
							</div>
						</div>
						<a href="/model/audi/1213/audi-a6-2014" title="">
							<img src="/resources/images/variant/1816/thumb_a6_15.jpg"
								 alt=""
								 class="thumb-image"/>
						</a>
					</div>
					<div class="auto-info-wrap">
						<div class="auto-naziv">
							<a href="/model/audi/1213/audi-a6-2014" title="2014 Audi A6 (C7 restyle)">
								2014 Audi A6 (C7 restyle)							</a>
						</div>
						<div class="stars">
							<div class="stars5"></div>
						</div>
						<div class="euro-ncap">
							EuroNCAP:&nbsp;
							91% occupant safety						</div>
						<div>
							<b>
								<a href="/model/audi/1213/audi-a6-2014" title="2014 Audi A6 (C7 restyle)" class="orange">
							   	   	more details...						   	   </a>
					   	   </b>
						</div>
					</div>
				</div>
							<div class="auto-wrap">
					<div class="auto-slika">
						<div class="auto-score-wrap clearfix">
														<div class="score-left">
								<div class="score5b">
									<p>drivers'</p>
									<span>4.5</span>
									<p>rating</p>
								</div>
							</div>
							<div class="score-right clearfix">
								<div class="score5a"></div>
								<div class="score4a"></div>
								<div class="score3a"></div>
								<div class="score2a"></div>
								<div class="score1a"></div>
							</div>
						</div>
						<a href="/model/audi/1424/audi-a6-2018" title="">
							<img src="/resources/images/variant/2161/thumb_2018_audi_a6.jpg"
								 alt=""
								 class="thumb-image"/>
						</a>
					</div>
					<div class="auto-info-wrap">
						<div class="auto-naziv">
							<a href="/model/audi/1424/audi-a6-2018" title="2018 Audi A6 (C8)">
								2018 Audi A6 (C8)							</a>
						</div>
						<div class="stars">
							<div class="stars5"></div>
						</div>
						<div class="euro-ncap">
							EuroNCAP:&nbsp;
							93% occupant safety						</div>
						<div>
							<b>
								<a href="/model/audi/1424/audi-a6-2018" title="2018 Audi A6 (C8)" class="orange">
							   	   	more details...						   	   </a>
					   	   </b>
						</div>
					</div>
				</div>
							<div class="auto-wrap">
					<div class="auto-slika">
						<div class="auto-score-wrap clearfix">
														<div class="score-left">
								<div class="score5b">
									<p>drivers'</p>
									<span>5.0</span>
									<p>rating</p>
								</div>
							</div>
							<div class="score-right clearfix">
								<div class="score5a"></div>
								<div class="score4a"></div>
								<div class="score3a"></div>
								<div class="score2a"></div>
								<div class="score1a"></div>
							</div>
						</div>
						<a href="/model/audi/233/audi-a7-2010" title="">
							<img src="/resources/images/variant/71/thumb_a7_1.jpg"
								 alt=""
								 class="thumb-image"/>
						</a>
					</div>
					<div class="auto-info-wrap">
						<div class="auto-naziv">
							<a href="/model/audi/233/audi-a7-2010" title="2010 Audi A7 (4G)">
								2010 Audi A7 (4G)							</a>
						</div>
						<div class="stars">
							<div class="stars"></div>
						</div>
						<div class="euro-ncap">
							EuroNCAP:&nbsp;
							not tested						</div>
						<div>
							<b>
								<a href="/model/audi/233/audi-a7-2010" title="2010 Audi A7 (4G)" class="orange">
							   	   	more details...						   	   </a>
					   	   </b>
						</div>
					</div>
				</div>
							<div class="auto-wrap">
					<div class="auto-slika">
						<div class="auto-score-wrap clearfix">
														<div class="score-left">
								<div class="score5b">
									<p>drivers'</p>
									<span>4.6</span>
									<p>rating</p>
								</div>
							</div>
							<div class="score-right clearfix">
								<div class="score5a"></div>
								<div class="score4a"></div>
								<div class="score3a"></div>
								<div class="score2a"></div>
								<div class="score1a"></div>
							</div>
						</div>
						<a href="/model/audi/234/audi-a7-2014" title="">
							<img src="/resources/images/variant/72/thumb_a7_2.jpg"
								 alt=""
								 class="thumb-image"/>
						</a>
					</div>
					<div class="auto-info-wrap">
						<div class="auto-naziv">
							<a href="/model/audi/234/audi-a7-2014" title="2014 Audi A7 (4G restyle)">
								2014 Audi A7 (4G restyle)							</a>
						</div>
						<div class="stars">
							<div class="stars"></div>
						</div>
						<div class="euro-ncap">
							EuroNCAP:&nbsp;
							not tested						</div>
						<div>
							<b>
								<a href="/model/audi/234/audi-a7-2014" title="2014 Audi A7 (4G restyle)" class="orange">
							   	   	more details...						   	   </a>
					   	   </b>
						</div>
					</div>
				</div>
							<div class="auto-wrap">
					<div class="auto-slika">
						<div class="auto-score-wrap clearfix">
														<div class="score-left">
								<div class="score3b">
									<p>drivers'</p>
									<span>3.0</span>
									<p>rating</p>
								</div>
							</div>
							<div class="score-right clearfix">
								<div class="score5a"></div>
								<div class="score4a"></div>
								<div class="score3a"></div>
								<div class="score2a"></div>
								<div class="score1a"></div>
							</div>
						</div>
						<a href="/model/audi/1425/audi-a7-2018" title="">
							<img src="/resources/images/variant/2167/thumb_2018_audi_a7.jpg"
								 alt=""
								 class="thumb-image"/>
						</a>
					</div>
					<div class="auto-info-wrap">
						<div class="auto-naziv">
							<a href="/model/audi/1425/audi-a7-2018" title="2018 Audi A7 (4K)">
								2018 Audi A7 (4K)							</a>
						</div>
						<div class="stars">
							<div class="stars5"></div>
						</div>
						<div class="euro-ncap">
							EuroNCAP:&nbsp;
							93% occupant safety						</div>
						<div>
							<b>
								<a href="/model/audi/1425/audi-a7-2018" title="2018 Audi A7 (4K)" class="orange">
							   	   	more details...						   	   </a>
					   	   </b>
						</div>
					</div>
				</div>
					</div>
</div>
<div class="proizvodnja-wrap">


		<div class="proizvodnja-box variant-header accordion-controller segment-header"
			 related-accordion="accordion-F"
			 segment="F"
		 >
			<div class="proizvodnja-godina">
				F - Executive car			</div>
			<div class="proizvodnja-model">
				6 models			</div>
		</div>
		<div id="accordion-F" class="variant-accordion ">
							<div class="auto-wrap">
					<div class="auto-slika">
						<div class="auto-score-wrap clearfix">
														<div class="score-left">
								<div class="score4b">
									<p>drivers'</p>
									<span>4.2</span>
									<p>rating</p>
								</div>
							</div>
							<div class="score-right clearfix">
								<div class="score5a"></div>
								<div class="score4a"></div>
								<div class="score3a"></div>
								<div class="score2a"></div>
								<div class="score1a"></div>
							</div>
						</div>
						<a href="/model/audi/235/audi-a8-1999" title="">
							<img src="/resources/images/variant/73/thumb_a8_1.jpg"
								 alt=""
								 class="thumb-image"/>
						</a>
					</div>
					<div class="auto-info-wrap">
						<div class="auto-naziv">
							<a href="/model/audi/235/audi-a8-1999" title="1999 Audi A8 (D2 restyle)">
								1999 Audi A8 (D2 restyle)							</a>
						</div>
						<div class="stars">
							<div class="stars"></div>
						</div>
						<div class="euro-ncap">
							EuroNCAP:&nbsp;
							not tested						</div>
						<div>
							<b>
								<a href="/model/audi/235/audi-a8-1999" title="1999 Audi A8 (D2 restyle)" class="orange">
							   	   	more details...						   	   </a>
					   	   </b>
						</div>
					</div>
				</div>
							<div class="auto-wrap">
					<div class="auto-slika">
						<div class="auto-score-wrap clearfix">
														<div class="score-left">
								<div class="score5b">
									<p>drivers'</p>
									<span>4.7</span>
									<p>rating</p>
								</div>
							</div>
							<div class="score-right clearfix">
								<div class="score5a"></div>
								<div class="score4a"></div>
								<div class="score3a"></div>
								<div class="score2a"></div>
								<div class="score1a"></div>
							</div>
						</div>
						<a href="/model/audi/236/audi-a8-2002" title="">
							<img src="/resources/images/variant/74/thumb_a8_2.jpg"
								 alt=""
								 class="thumb-image"/>
						</a>
					</div>
					<div class="auto-info-wrap">
						<div class="auto-naziv">
							<a href="/model/audi/236/audi-a8-2002" title="2002 Audi A8 (D3)">
								2002 Audi A8 (D3)							</a>
						</div>
						<div class="stars">
							<div class="stars"></div>
						</div>
						<div class="euro-ncap">
							EuroNCAP:&nbsp;
							not tested						</div>
						<div>
							<b>
								<a href="/model/audi/236/audi-a8-2002" title="2002 Audi A8 (D3)" class="orange">
							   	   	more details...						   	   </a>
					   	   </b>
						</div>
					</div>
				</div>
							<div class="auto-wrap">
					<div class="auto-slika">
						<div class="auto-score-wrap clearfix">
														<div class="score-left">
								<div class="score4b">
									<p>drivers'</p>
									<span>4.3</span>
									<p>rating</p>
								</div>
							</div>
							<div class="score-right clearfix">
								<div class="score5a"></div>
								<div class="score4a"></div>
								<div class="score3a"></div>
								<div class="score2a"></div>
								<div class="score1a"></div>
							</div>
						</div>
						<a href="/model/audi/237/audi-a8-2007" title="">
							<img src="/resources/images/variant/75/thumb_a8_3.jpg"
								 alt=""
								 class="thumb-image"/>
						</a>
					</div>
					<div class="auto-info-wrap">
						<div class="auto-naziv">
							<a href="/model/audi/237/audi-a8-2007" title="2007 Audi A8 (D3 restyle)">
								2007 Audi A8 (D3 restyle)							</a>
						</div>
						<div class="stars">
							<div class="stars"></div>
						</div>
						<div class="euro-ncap">
							EuroNCAP:&nbsp;
							not tested						</div>
						<div>
							<b>
								<a href="/model/audi/237/audi-a8-2007" title="2007 Audi A8 (D3 restyle)" class="orange">
							   	   	more details...						   	   </a>
					   	   </b>
						</div>
					</div>
				</div>
							<div class="auto-wrap">
					<div class="auto-slika">
						<div class="auto-score-wrap clearfix">
														<div class="score-left">
								<div class="score5b">
									<p>drivers'</p>
									<span>4.7</span>
									<p>rating</p>
								</div>
							</div>
							<div class="score-right clearfix">
								<div class="score5a"></div>
								<div class="score4a"></div>
								<div class="score3a"></div>
								<div class="score2a"></div>
								<div class="score1a"></div>
							</div>
						</div>
						<a href="/model/audi/238/audi-a8-2010" title="">
							<img src="/resources/images/variant/76/thumb_a8_4.jpg"
								 alt=""
								 class="thumb-image"/>
						</a>
					</div>
					<div class="auto-info-wrap">
						<div class="auto-naziv">
							<a href="/model/audi/238/audi-a8-2010" title="2010 Audi A8 (D4)">
								2010 Audi A8 (D4)							</a>
						</div>
						<div class="stars">
							<div class="stars"></div>
						</div>
						<div class="euro-ncap">
							EuroNCAP:&nbsp;
							not tested						</div>
						<div>
							<b>
								<a href="/model/audi/238/audi-a8-2010" title="2010 Audi A8 (D4)" class="orange">
							   	   	more details...						   	   </a>
					   	   </b>
						</div>
					</div>
				</div>
							<div class="auto-wrap">
					<div class="auto-slika">
						<div class="auto-score-wrap clearfix">
														<div class="score-left">
								<div class="score5b">
									<p>drivers'</p>
									<span>5.0</span>
									<p>rating</p>
								</div>
							</div>
							<div class="score-right clearfix">
								<div class="score5a"></div>
								<div class="score4a"></div>
								<div class="score3a"></div>
								<div class="score2a"></div>
								<div class="score1a"></div>
							</div>
						</div>
						<a href="/model/audi/239/audi-a8-2013" title="">
							<img src="/resources/images/variant/77/thumb_a8_5.jpg"
								 alt=""
								 class="thumb-image"/>
						</a>
					</div>
					<div class="auto-info-wrap">
						<div class="auto-naziv">
							<a href="/model/audi/239/audi-a8-2013" title="2013 Audi A8 (D4 restyle)">
								2013 Audi A8 (D4 restyle)							</a>
						</div>
						<div class="stars">
							<div class="stars"></div>
						</div>
						<div class="euro-ncap">
							EuroNCAP:&nbsp;
							not tested						</div>
						<div>
							<b>
								<a href="/model/audi/239/audi-a8-2013" title="2013 Audi A8 (D4 restyle)" class="orange">
							   	   	more details...						   	   </a>
					   	   </b>
						</div>
					</div>
				</div>
							<div class="auto-wrap">
					<div class="auto-slika">
						<div class="auto-score-wrap clearfix">
														<div class="score-left">
								<div class="score3b">
									<p>drivers'</p>
									<span>3.0</span>
									<p>rating</p>
								</div>
							</div>
							<div class="score-right clearfix">
								<div class="score5a"></div>
								<div class="score4a"></div>
								<div class="score3a"></div>
								<div class="score2a"></div>
								<div class="score1a"></div>
							</div>
						</div>
						<a href="/model/audi/1454/audi-a8-2017" title="">
							<img src="/resources/images/variant/2205/thumb_2017_audi_a8.jpg"
								 alt=""
								 class="thumb-image"/>
						</a>
					</div>
					<div class="auto-info-wrap">
						<div class="auto-naziv">
							<a href="/model/audi/1454/audi-a8-2017" title="2017 Audi A8 (D5)">
								2017 Audi A8 (D5)							</a>
						</div>
						<div class="stars">
							<div class="stars"></div>
						</div>
						<div class="euro-ncap">
							EuroNCAP:&nbsp;
							not tested						</div>
						<div>
							<b>
								<a href="/model/audi/1454/audi-a8-2017" title="2017 Audi A8 (D5)" class="orange">
							   	   	more details...						   	   </a>
					   	   </b>
						</div>
					</div>
				</div>
					</div>
</div>
<div class="proizvodnja-wrap">


		<div class="proizvodnja-box variant-header accordion-controller segment-header"
			 related-accordion="accordion-S"
			 segment="S"
		 >
			<div class="proizvodnja-godina">
				S - Sports car			</div>
			<div class="proizvodnja-model">
				9 models			</div>
		</div>
		<div id="accordion-S" class="variant-accordion ">
							<div class="auto-wrap">
					<div class="auto-slika">
						<div class="auto-score-wrap clearfix">
														<div class="score-left">
								<div class="score4b">
									<p>drivers'</p>
									<span>4.5</span>
									<p>rating</p>
								</div>
							</div>
							<div class="score-right clearfix">
								<div class="score5a"></div>
								<div class="score4a"></div>
								<div class="score3a"></div>
								<div class="score2a"></div>
								<div class="score1a"></div>
							</div>
						</div>
						<a href="/model/audi/1053/audi-a5-coupe-2007" title="">
							<img src="/resources/images/variant/49/thumb_a5_1.jpg"
								 alt=""
								 class="thumb-image"/>
						</a>
					</div>
					<div class="auto-info-wrap">
						<div class="auto-naziv">
							<a href="/model/audi/1053/audi-a5-coupe-2007" title="2007 Audi A5 Coupe (8T)">
								2007 Audi A5 Coupe (8T)							</a>
						</div>
						<div class="stars">
							<div class="stars"></div>
						</div>
						<div class="euro-ncap">
							EuroNCAP:&nbsp;
							not tested						</div>
						<div>
							<b>
								<a href="/model/audi/1053/audi-a5-coupe-2007" title="2007 Audi A5 Coupe (8T)" class="orange">
							   	   	more details...						   	   </a>
					   	   </b>
						</div>
					</div>
				</div>
							<div class="auto-wrap">
					<div class="auto-slika">
						<div class="auto-score-wrap clearfix">
														<div class="score-left">
								<div class="score5b">
									<p>drivers'</p>
									<span>4.8</span>
									<p>rating</p>
								</div>
							</div>
							<div class="score-right clearfix">
								<div class="score5a"></div>
								<div class="score4a"></div>
								<div class="score3a"></div>
								<div class="score2a"></div>
								<div class="score1a"></div>
							</div>
						</div>
						<a href="/model/audi/1054/audi-a5-coupe-2011" title="">
							<img src="/resources/images/variant/51/thumb_a5_3.jpg"
								 alt=""
								 class="thumb-image"/>
						</a>
					</div>
					<div class="auto-info-wrap">
						<div class="auto-naziv">
							<a href="/model/audi/1054/audi-a5-coupe-2011" title="2011 Audi A5 Coupe (8T restyle)">
								2011 Audi A5 Coupe (8T restyle)							</a>
						</div>
						<div class="stars">
							<div class="stars"></div>
						</div>
						<div class="euro-ncap">
							EuroNCAP:&nbsp;
							not tested						</div>
						<div>
							<b>
								<a href="/model/audi/1054/audi-a5-coupe-2011" title="2011 Audi A5 Coupe (8T restyle)" class="orange">
							   	   	more details...						   	   </a>
					   	   </b>
						</div>
					</div>
				</div>
							<div class="auto-wrap">
					<div class="auto-slika">
						<div class="auto-score-wrap clearfix">
														<div class="score-left">
								<div class="score3b">
									<p>drivers'</p>
									<span>3.2</span>
									<p>rating</p>
								</div>
							</div>
							<div class="score-right clearfix">
								<div class="score5a"></div>
								<div class="score4a"></div>
								<div class="score3a"></div>
								<div class="score2a"></div>
								<div class="score1a"></div>
							</div>
						</div>
						<a href="/model/audi/1246/audi-a5-coupe-2016" title="">
							<img src="/resources/images/variant/1879/thumb_a5_8.jpg"
								 alt=""
								 class="thumb-image"/>
						</a>
					</div>
					<div class="auto-info-wrap">
						<div class="auto-naziv">
							<a href="/model/audi/1246/audi-a5-coupe-2016" title="2016 Audi A5 Coupe (F5)">
								2016 Audi A5 Coupe (F5)							</a>
						</div>
						<div class="stars">
							<div class="stars5"></div>
						</div>
						<div class="euro-ncap">
							EuroNCAP:&nbsp;
							89% occupant safety						</div>
						<div>
							<b>
								<a href="/model/audi/1246/audi-a5-coupe-2016" title="2016 Audi A5 Coupe (F5)" class="orange">
							   	   	more details...						   	   </a>
					   	   </b>
						</div>
					</div>
				</div>
							<div class="auto-wrap">
					<div class="auto-slika">
						<div class="auto-score-wrap clearfix">
														<div class="score-left">
								<div class="score3b">
									<p>drivers'</p>
									<span>3.0</span>
									<p>rating</p>
								</div>
							</div>
							<div class="score-right clearfix">
								<div class="score5a"></div>
								<div class="score4a"></div>
								<div class="score3a"></div>
								<div class="score2a"></div>
								<div class="score1a"></div>
							</div>
						</div>
						<a href="/model/audi/1340/audi-r8-2007" title="">
							<img src="/resources/images/variant/2039/thumb_r8_1.jpg"
								 alt=""
								 class="thumb-image"/>
						</a>
					</div>
					<div class="auto-info-wrap">
						<div class="auto-naziv">
							<a href="/model/audi/1340/audi-r8-2007" title="2007 Audi R8 (Typ 42)">
								2007 Audi R8 (Typ 42)							</a>
						</div>
						<div class="stars">
							<div class="stars"></div>
						</div>
						<div class="euro-ncap">
							EuroNCAP:&nbsp;
							not tested						</div>
						<div>
							<b>
								<a href="/model/audi/1340/audi-r8-2007" title="2007 Audi R8 (Typ 42)" class="orange">
							   	   	more details...						   	   </a>
					   	   </b>
						</div>
					</div>
				</div>
							<div class="auto-wrap">
					<div class="auto-slika">
						<div class="auto-score-wrap clearfix">
														<div class="score-left">
								<div class="score3b">
									<p>drivers'</p>
									<span>3.0</span>
									<p>rating</p>
								</div>
							</div>
							<div class="score-right clearfix">
								<div class="score5a"></div>
								<div class="score4a"></div>
								<div class="score3a"></div>
								<div class="score2a"></div>
								<div class="score1a"></div>
							</div>
						</div>
						<a href="/model/audi/1341/audi-r8-2012" title="">
							<img src="/resources/images/variant/2041/thumb_r8_3.jpg"
								 alt=""
								 class="thumb-image"/>
						</a>
					</div>
					<div class="auto-info-wrap">
						<div class="auto-naziv">
							<a href="/model/audi/1341/audi-r8-2012" title="2012 Audi R8 (Typ 42 restyle)">
								2012 Audi R8 (Typ 42 restyle)							</a>
						</div>
						<div class="stars">
							<div class="stars"></div>
						</div>
						<div class="euro-ncap">
							EuroNCAP:&nbsp;
							not tested						</div>
						<div>
							<b>
								<a href="/model/audi/1341/audi-r8-2012" title="2012 Audi R8 (Typ 42 restyle)" class="orange">
							   	   	more details...						   	   </a>
					   	   </b>
						</div>
					</div>
				</div>
							<div class="auto-wrap">
					<div class="auto-slika">
						<div class="auto-score-wrap clearfix">
														<div class="score-left">
								<div class="score4b">
									<p>drivers'</p>
									<span>4.0</span>
									<p>rating</p>
								</div>
							</div>
							<div class="score-right clearfix">
								<div class="score5a"></div>
								<div class="score4a"></div>
								<div class="score3a"></div>
								<div class="score2a"></div>
								<div class="score1a"></div>
							</div>
						</div>
						<a href="/model/audi/240/audi-tt-1998" title="">
							<img src="/resources/images/variant/84/thumb_tt_1.jpg"
								 alt=""
								 class="thumb-image"/>
						</a>
					</div>
					<div class="auto-info-wrap">
						<div class="auto-naziv">
							<a href="/model/audi/240/audi-tt-1998" title="1998 Audi TT (8N)">
								1998 Audi TT (8N)							</a>
						</div>
						<div class="stars">
							<div class="stars4"></div>
						</div>
						<div class="euro-ncap">
							EuroNCAP:&nbsp;
							~70% occupant safety						</div>
						<div>
							<b>
								<a href="/model/audi/240/audi-tt-1998" title="1998 Audi TT (8N)" class="orange">
							   	   	more details...						   	   </a>
					   	   </b>
						</div>
					</div>
				</div>
							<div class="auto-wrap">
					<div class="auto-slika">
						<div class="auto-score-wrap clearfix">
														<div class="score-left">
								<div class="score5b">
									<p>drivers'</p>
									<span>4.6</span>
									<p>rating</p>
								</div>
							</div>
							<div class="score-right clearfix">
								<div class="score5a"></div>
								<div class="score4a"></div>
								<div class="score3a"></div>
								<div class="score2a"></div>
								<div class="score1a"></div>
							</div>
						</div>
						<a href="/model/audi/241/audi-tt-2006" title="">
							<img src="/resources/images/variant/86/thumb_tt_3.jpg"
								 alt=""
								 class="thumb-image"/>
						</a>
					</div>
					<div class="auto-info-wrap">
						<div class="auto-naziv">
							<a href="/model/audi/241/audi-tt-2006" title="2006 Audi TT (8J)">
								2006 Audi TT (8J)							</a>
						</div>
						<div class="stars">
							<div class="stars4"></div>
						</div>
						<div class="euro-ncap">
							EuroNCAP:&nbsp;
							~70% occupant safety						</div>
						<div>
							<b>
								<a href="/model/audi/241/audi-tt-2006" title="2006 Audi TT (8J)" class="orange">
							   	   	more details...						   	   </a>
					   	   </b>
						</div>
					</div>
				</div>
							<div class="auto-wrap">
					<div class="auto-slika">
						<div class="auto-score-wrap clearfix">
														<div class="score-left">
								<div class="score5b">
									<p>drivers'</p>
									<span>5.0</span>
									<p>rating</p>
								</div>
							</div>
							<div class="score-right clearfix">
								<div class="score5a"></div>
								<div class="score4a"></div>
								<div class="score3a"></div>
								<div class="score2a"></div>
								<div class="score1a"></div>
							</div>
						</div>
						<a href="/model/audi/242/audi-tt-2010" title="">
							<img src="/resources/images/variant/88/thumb_tt_5.jpg"
								 alt=""
								 class="thumb-image"/>
						</a>
					</div>
					<div class="auto-info-wrap">
						<div class="auto-naziv">
							<a href="/model/audi/242/audi-tt-2010" title="2010 Audi TT (8J restyle)">
								2010 Audi TT (8J restyle)							</a>
						</div>
						<div class="stars">
							<div class="stars4"></div>
						</div>
						<div class="euro-ncap">
							EuroNCAP:&nbsp;
							~70% occupant safety						</div>
						<div>
							<b>
								<a href="/model/audi/242/audi-tt-2010" title="2010 Audi TT (8J restyle)" class="orange">
							   	   	more details...						   	   </a>
					   	   </b>
						</div>
					</div>
				</div>
							<div class="auto-wrap">
					<div class="auto-slika">
						<div class="auto-score-wrap clearfix">
														<div class="score-left">
								<div class="score5b">
									<p>drivers'</p>
									<span>5.0</span>
									<p>rating</p>
								</div>
							</div>
							<div class="score-right clearfix">
								<div class="score5a"></div>
								<div class="score4a"></div>
								<div class="score3a"></div>
								<div class="score2a"></div>
								<div class="score1a"></div>
							</div>
						</div>
						<a href="/model/audi/1057/audi-tt-2015" title="">
							<img src="/resources/images/variant/280/thumb_tt_7.jpg"
								 alt=""
								 class="thumb-image"/>
						</a>
					</div>
					<div class="auto-info-wrap">
						<div class="auto-naziv">
							<a href="/model/audi/1057/audi-tt-2015" title="2015 Audi TT (FV)">
								2015 Audi TT (FV)							</a>
						</div>
						<div class="stars">
							<div class="stars4"></div>
						</div>
						<div class="euro-ncap">
							EuroNCAP:&nbsp;
							81% occupant safety						</div>
						<div>
							<b>
								<a href="/model/audi/1057/audi-tt-2015" title="2015 Audi TT (FV)" class="orange">
							   	   	more details...						   	   </a>
					   	   </b>
						</div>
					</div>
				</div>
					</div>
</div>
<div class="proizvodnja-wrap">


		<div class="proizvodnja-box variant-header accordion-controller segment-header"
			 related-accordion="accordion-M"
			 segment="M"
		 >
			<div class="proizvodnja-godina">
				M - MPV			</div>
			<div class="proizvodnja-model">
				1 model			</div>
		</div>
		<div id="accordion-M" class="variant-accordion ">
							<div class="auto-wrap">
					<div class="auto-slika">
						<div class="auto-score-wrap clearfix">
														<div class="score-left">
								<div class="score4b">
									<p>drivers'</p>
									<span>4.3</span>
									<p>rating</p>
								</div>
							</div>
							<div class="score-right clearfix">
								<div class="score5a"></div>
								<div class="score4a"></div>
								<div class="score3a"></div>
								<div class="score2a"></div>
								<div class="score1a"></div>
							</div>
						</div>
						<a href="/model/audi/213/audi-a2-2000" title="">
							<img src="/resources/images/variant/24/thumb_a2.jpg"
								 alt=""
								 class="thumb-image"/>
						</a>
					</div>
					<div class="auto-info-wrap">
						<div class="auto-naziv">
							<a href="/model/audi/213/audi-a2-2000" title="2000 Audi A2 (8Z)">
								2000 Audi A2 (8Z)							</a>
						</div>
						<div class="stars">
							<div class="stars4"></div>
						</div>
						<div class="euro-ncap">
							EuroNCAP:&nbsp;
							~70% occupant safety						</div>
						<div>
							<b>
								<a href="/model/audi/213/audi-a2-2000" title="2000 Audi A2 (8Z)" class="orange">
							   	   	more details...						   	   </a>
					   	   </b>
						</div>
					</div>
				</div>
					</div>
</div>
<div class="proizvodnja-wrap">


		<div class="proizvodnja-box variant-header accordion-controller segment-header"
			 related-accordion="accordion-J"
			 segment="J"
		 >
			<div class="proizvodnja-godina">
				J - SUV			</div>
			<div class="proizvodnja-model">
				11 models			</div>
		</div>
		<div id="accordion-J" class="variant-accordion ">
							<div class="auto-wrap">
					<div class="auto-slika">
						<div class="auto-score-wrap clearfix">
														<div class="score-left">
								<div class="score3b">
									<p>drivers'</p>
									<span>3.0</span>
									<p>rating</p>
								</div>
							</div>
							<div class="score-right clearfix">
								<div class="score5a"></div>
								<div class="score4a"></div>
								<div class="score3a"></div>
								<div class="score2a"></div>
								<div class="score1a"></div>
							</div>
						</div>
						<a href="/model/audi/1291/audi-q2-2016" title="">
							<img src="/resources/images/variant/1949/thumb_q2_1.jpg"
								 alt=""
								 class="thumb-image"/>
						</a>
					</div>
					<div class="auto-info-wrap">
						<div class="auto-naziv">
							<a href="/model/audi/1291/audi-q2-2016" title="2016 Audi Q2 (GA)">
								2016 Audi Q2 (GA)							</a>
						</div>
						<div class="stars">
							<div class="stars5"></div>
						</div>
						<div class="euro-ncap">
							EuroNCAP:&nbsp;
							93% occupant safety						</div>
						<div>
							<b>
								<a href="/model/audi/1291/audi-q2-2016" title="2016 Audi Q2 (GA)" class="orange">
							   	   	more details...						   	   </a>
					   	   </b>
						</div>
					</div>
				</div>
							<div class="auto-wrap">
					<div class="auto-slika">
						<div class="auto-score-wrap clearfix">
														<div class="score-left">
								<div class="score5b">
									<p>drivers'</p>
									<span>5.0</span>
									<p>rating</p>
								</div>
							</div>
							<div class="score-right clearfix">
								<div class="score5a"></div>
								<div class="score4a"></div>
								<div class="score3a"></div>
								<div class="score2a"></div>
								<div class="score1a"></div>
							</div>
						</div>
						<a href="/model/audi/243/audi-q3-2011" title="">
							<img src="/resources/images/variant/78/thumb_q3_1.jpg"
								 alt=""
								 class="thumb-image"/>
						</a>
					</div>
					<div class="auto-info-wrap">
						<div class="auto-naziv">
							<a href="/model/audi/243/audi-q3-2011" title="2011 Audi Q3 (8U)">
								2011 Audi Q3 (8U)							</a>
						</div>
						<div class="stars">
							<div class="stars5"></div>
						</div>
						<div class="euro-ncap">
							EuroNCAP:&nbsp;
							94% occupant safety						</div>
						<div>
							<b>
								<a href="/model/audi/243/audi-q3-2011" title="2011 Audi Q3 (8U)" class="orange">
							   	   	more details...						   	   </a>
					   	   </b>
						</div>
					</div>
				</div>
							<div class="auto-wrap">
					<div class="auto-slika">
						<div class="auto-score-wrap clearfix">
														<div class="score-left">
								<div class="score5b">
									<p>drivers'</p>
									<span>5.0</span>
									<p>rating</p>
								</div>
							</div>
							<div class="score-right clearfix">
								<div class="score5a"></div>
								<div class="score4a"></div>
								<div class="score3a"></div>
								<div class="score2a"></div>
								<div class="score1a"></div>
							</div>
						</div>
						<a href="/model/audi/1058/audi-q3-2015" title="">
							<img src="/resources/images/variant/281/thumb_q3_2.jpg"
								 alt=""
								 class="thumb-image"/>
						</a>
					</div>
					<div class="auto-info-wrap">
						<div class="auto-naziv">
							<a href="/model/audi/1058/audi-q3-2015" title="2015 Audi Q3 (8U restyle)">
								2015 Audi Q3 (8U restyle)							</a>
						</div>
						<div class="stars">
							<div class="stars5"></div>
						</div>
						<div class="euro-ncap">
							EuroNCAP:&nbsp;
							94% occupant safety						</div>
						<div>
							<b>
								<a href="/model/audi/1058/audi-q3-2015" title="2015 Audi Q3 (8U restyle)" class="orange">
							   	   	more details...						   	   </a>
					   	   </b>
						</div>
					</div>
				</div>
							<div class="auto-wrap">
					<div class="auto-slika">
						<div class="auto-score-wrap clearfix">
														<div class="score-left">
								<div class="score3b">
									<p>drivers'</p>
									<span>3.0</span>
									<p>rating</p>
								</div>
							</div>
							<div class="score-right clearfix">
								<div class="score5a"></div>
								<div class="score4a"></div>
								<div class="score3a"></div>
								<div class="score2a"></div>
								<div class="score1a"></div>
							</div>
						</div>
						<a href="/model/audi/1414/audi-q3-2018" title="">
							<img src="/resources/images/variant/2144/thumb_2019_audi_q3.jpg"
								 alt=""
								 class="thumb-image"/>
						</a>
					</div>
					<div class="auto-info-wrap">
						<div class="auto-naziv">
							<a href="/model/audi/1414/audi-q3-2018" title="2018 Audi Q3 (F3)">
								2018 Audi Q3 (F3)							</a>
						</div>
						<div class="stars">
							<div class="stars5"></div>
						</div>
						<div class="euro-ncap">
							EuroNCAP:&nbsp;
							95% occupant safety						</div>
						<div>
							<b>
								<a href="/model/audi/1414/audi-q3-2018" title="2018 Audi Q3 (F3)" class="orange">
							   	   	more details...						   	   </a>
					   	   </b>
						</div>
					</div>
				</div>
							<div class="auto-wrap">
					<div class="auto-slika">
						<div class="auto-score-wrap clearfix">
														<div class="score-left">
								<div class="score5b">
									<p>drivers'</p>
									<span>5.0</span>
									<p>rating</p>
								</div>
							</div>
							<div class="score-right clearfix">
								<div class="score5a"></div>
								<div class="score4a"></div>
								<div class="score3a"></div>
								<div class="score2a"></div>
								<div class="score1a"></div>
							</div>
						</div>
						<a href="/model/audi/244/audi-q5-2008" title="">
							<img src="/resources/images/variant/79/thumb_q5_1.jpg"
								 alt=""
								 class="thumb-image"/>
						</a>
					</div>
					<div class="auto-info-wrap">
						<div class="auto-naziv">
							<a href="/model/audi/244/audi-q5-2008" title="2008 Audi Q5 (8R)">
								2008 Audi Q5 (8R)							</a>
						</div>
						<div class="stars">
							<div class="stars5"></div>
						</div>
						<div class="euro-ncap">
							EuroNCAP:&nbsp;
							92% occupant safety						</div>
						<div>
							<b>
								<a href="/model/audi/244/audi-q5-2008" title="2008 Audi Q5 (8R)" class="orange">
							   	   	more details...						   	   </a>
					   	   </b>
						</div>
					</div>
				</div>
							<div class="auto-wrap">
					<div class="auto-slika">
						<div class="auto-score-wrap clearfix">
														<div class="score-left">
								<div class="score4b">
									<p>drivers'</p>
									<span>4.3</span>
									<p>rating</p>
								</div>
							</div>
							<div class="score-right clearfix">
								<div class="score5a"></div>
								<div class="score4a"></div>
								<div class="score3a"></div>
								<div class="score2a"></div>
								<div class="score1a"></div>
							</div>
						</div>
						<a href="/model/audi/245/audi-q5-2012" title="">
							<img src="/resources/images/variant/80/thumb_q5_2.jpg"
								 alt=""
								 class="thumb-image"/>
						</a>
					</div>
					<div class="auto-info-wrap">
						<div class="auto-naziv">
							<a href="/model/audi/245/audi-q5-2012" title="2012 Audi Q5 (8R restyle)">
								2012 Audi Q5 (8R restyle)							</a>
						</div>
						<div class="stars">
							<div class="stars5"></div>
						</div>
						<div class="euro-ncap">
							EuroNCAP:&nbsp;
							92% occupant safety						</div>
						<div>
							<b>
								<a href="/model/audi/245/audi-q5-2012" title="2012 Audi Q5 (8R restyle)" class="orange">
							   	   	more details...						   	   </a>
					   	   </b>
						</div>
					</div>
				</div>
							<div class="auto-wrap">
					<div class="auto-slika">
						<div class="auto-score-wrap clearfix">
														<div class="score-left">
								<div class="score5b">
									<p>drivers'</p>
									<span>4.5</span>
									<p>rating</p>
								</div>
							</div>
							<div class="score-right clearfix">
								<div class="score5a"></div>
								<div class="score4a"></div>
								<div class="score3a"></div>
								<div class="score2a"></div>
								<div class="score1a"></div>
							</div>
						</div>
						<a href="/model/audi/1304/audi-q5-2016" title="">
							<img src="/resources/images/variant/1972/thumb_q5_3.jpg"
								 alt=""
								 class="thumb-image"/>
						</a>
					</div>
					<div class="auto-info-wrap">
						<div class="auto-naziv">
							<a href="/model/audi/1304/audi-q5-2016" title="2016 Audi Q5 (FY)">
								2016 Audi Q5 (FY)							</a>
						</div>
						<div class="stars">
							<div class="stars5"></div>
						</div>
						<div class="euro-ncap">
							EuroNCAP:&nbsp;
							93% occupant safety						</div>
						<div>
							<b>
								<a href="/model/audi/1304/audi-q5-2016" title="2016 Audi Q5 (FY)" class="orange">
							   	   	more details...						   	   </a>
					   	   </b>
						</div>
					</div>
				</div>
							<div class="auto-wrap">
					<div class="auto-slika">
						<div class="auto-score-wrap clearfix">
														<div class="score-left">
								<div class="score5b">
									<p>drivers'</p>
									<span>4.9</span>
									<p>rating</p>
								</div>
							</div>
							<div class="score-right clearfix">
								<div class="score5a"></div>
								<div class="score4a"></div>
								<div class="score3a"></div>
								<div class="score2a"></div>
								<div class="score1a"></div>
							</div>
						</div>
						<a href="/model/audi/246/audi-q7-2006" title="">
							<img src="/resources/images/variant/81/thumb_q7_1.jpg"
								 alt=""
								 class="thumb-image"/>
						</a>
					</div>
					<div class="auto-info-wrap">
						<div class="auto-naziv">
							<a href="/model/audi/246/audi-q7-2006" title="2006 Audi Q7 (4L)">
								2006 Audi Q7 (4L)							</a>
						</div>
						<div class="stars">
							<div class="stars4"></div>
						</div>
						<div class="euro-ncap">
							EuroNCAP:&nbsp;
							~70% occupant safety						</div>
						<div>
							<b>
								<a href="/model/audi/246/audi-q7-2006" title="2006 Audi Q7 (4L)" class="orange">
							   	   	more details...						   	   </a>
					   	   </b>
						</div>
					</div>
				</div>
							<div class="auto-wrap">
					<div class="auto-slika">
						<div class="auto-score-wrap clearfix">
														<div class="score-left">
								<div class="score5b">
									<p>drivers'</p>
									<span>5.0</span>
									<p>rating</p>
								</div>
							</div>
							<div class="score-right clearfix">
								<div class="score5a"></div>
								<div class="score4a"></div>
								<div class="score3a"></div>
								<div class="score2a"></div>
								<div class="score1a"></div>
							</div>
						</div>
						<a href="/model/audi/247/audi-q7-2009" title="">
							<img src="/resources/images/variant/82/thumb_q7_2.jpg"
								 alt=""
								 class="thumb-image"/>
						</a>
					</div>
					<div class="auto-info-wrap">
						<div class="auto-naziv">
							<a href="/model/audi/247/audi-q7-2009" title="2009 Audi Q7 (4L restyle)">
								2009 Audi Q7 (4L restyle)							</a>
						</div>
						<div class="stars">
							<div class="stars4"></div>
						</div>
						<div class="euro-ncap">
							EuroNCAP:&nbsp;
							~70% occupant safety						</div>
						<div>
							<b>
								<a href="/model/audi/247/audi-q7-2009" title="2009 Audi Q7 (4L restyle)" class="orange">
							   	   	more details...						   	   </a>
					   	   </b>
						</div>
					</div>
				</div>
							<div class="auto-wrap">
					<div class="auto-slika">
						<div class="auto-score-wrap clearfix">
														<div class="score-left">
								<div class="score5b">
									<p>drivers'</p>
									<span>4.7</span>
									<p>rating</p>
								</div>
							</div>
							<div class="score-right clearfix">
								<div class="score5a"></div>
								<div class="score4a"></div>
								<div class="score3a"></div>
								<div class="score2a"></div>
								<div class="score1a"></div>
							</div>
						</div>
						<a href="/model/audi/1055/audi-q7-2015" title="">
							<img src="/resources/images/variant/83/thumb_q7_3.jpg"
								 alt=""
								 class="thumb-image"/>
						</a>
					</div>
					<div class="auto-info-wrap">
						<div class="auto-naziv">
							<a href="/model/audi/1055/audi-q7-2015" title="2015 Audi Q7 (4M)">
								2015 Audi Q7 (4M)							</a>
						</div>
						<div class="stars">
							<div class="stars5"></div>
						</div>
						<div class="euro-ncap">
							EuroNCAP:&nbsp;
							94% occupant safety						</div>
						<div>
							<b>
								<a href="/model/audi/1055/audi-q7-2015" title="2015 Audi Q7 (4M)" class="orange">
							   	   	more details...						   	   </a>
					   	   </b>
						</div>
					</div>
				</div>
							<div class="auto-wrap">
					<div class="auto-slika">
						<div class="auto-score-wrap clearfix">
														<div class="score-left">
								<div class="score3b">
									<p>drivers'</p>
									<span>3.0</span>
									<p>rating</p>
								</div>
							</div>
							<div class="score-right clearfix">
								<div class="score5a"></div>
								<div class="score4a"></div>
								<div class="score3a"></div>
								<div class="score2a"></div>
								<div class="score1a"></div>
							</div>
						</div>
						<a href="/model/audi/1430/audi-q8-2018" title="">
							<img src="/resources/images/variant/2175/thumb_2018_audi_q8.jpg"
								 alt=""
								 class="thumb-image"/>
						</a>
					</div>
					<div class="auto-info-wrap">
						<div class="auto-naziv">
							<a href="/model/audi/1430/audi-q8-2018" title="2018 Audi Q8 (4M)">
								2018 Audi Q8 (4M)							</a>
						</div>
						<div class="stars">
							<div class="stars5"></div>
						</div>
						<div class="euro-ncap">
							EuroNCAP:&nbsp;
							93% occupant safety						</div>
						<div>
							<b>
								<a href="/model/audi/1430/audi-q8-2018" title="2018 Audi Q8 (4M)" class="orange">
							   	   	more details...						   	   </a>
					   	   </b>
						</div>
					</div>
				</div>
					</div>
</div>

</div>
<div>
	<h2><i class="fa fa-search" aria-hidden="true"></i>Check a car by its VIN number</h2>
</div>
<div class="box-wrap" style="height: 250px;">
	<div
	  data-cvaff
	  data-locale="en-EU"
	  data-a="automanijak"
	  data-b="81ec5429"
	  style="width: 100%; height: 100%">
	</div>
</div>
	</div>
			</div>
		</div>
		<div id="footer">
			<div class="footer-wrap">
				Copyright &copy; 2015 - 2022 automaniac.org - All rights reserved.
				</br>
				Powered by <strong>Gama1 Solutions</strong>
				<div class="footer-links">
					<div class="footer-menu">
						<ul>
							<li><a href="/">HOME page</a></li>
							<li><a href="/about">about</a></li>
							<li><a href="/terms">terms & conditions</a></li>
							<li><a href="/legal">legal note</a></li>
							<li><a href="/faq">frequent questions (FAQ)</a></li>
							<li><a href="/contact">contact us</a></li>
						</ul>
					</div>

					<div class='footer-auto-brend'><ul><li><a href='/make/alfa-romeo'>Alfa Romeo</a></li><li><a href='/make/audi'>Audi</a></li><li><a href='/make/bmw'>BMW</a></li><li><a href='/make/chery'>Chery</a></li><li><a href='/make/chevrolet'>Chevrolet</a></li><li><a href='/make/chrysler'>Chrysler</a></li></ul></div><div class='footer-auto-brend'><ul><li><a href='/make/citroen'>Citroen</a></li><li><a href='/make/dacia'>Dacia</a></li><li><a href='/make/daihatsu'>Daihatsu</a></li><li><a href='/make/dodge'>Dodge</a></li><li><a href='/make/fiat'>FIAT</a></li><li><a href='/make/ford'>Ford</a></li></ul></div><div class='footer-auto-brend'><ul><li><a href='/make/honda'>Honda</a></li><li><a href='/make/hyundai'>Hyundai</a></li><li><a href='/make/infiniti'>Infiniti</a></li><li><a href='/make/jaguar'>Jaguar</a></li><li><a href='/make/jeep'>Jeep</a></li><li><a href='/make/kia'>KIA</a></li></ul></div><div class='footer-auto-brend'><ul><li><a href='/make/lada'>Lada</a></li><li><a href='/make/lancia'>Lancia</a></li><li><a href='/make/land-rover'>Land Rover</a></li><li><a href='/make/lexus'>Lexus</a></li><li><a href='/make/mazda'>Mazda</a></li><li><a href='/make/mercedes-benz'>Mercedes</a></li></ul></div><div class='footer-auto-brend'><ul><li><a href='/make/mini'>Mini</a></li><li><a href='/make/mitsubishi'>Mitsubishi</a></li><li><a href='/make/nissan'>Nissan</a></li><li><a href='/make/opel'>Opel</a></li><li><a href='/make/peugeot'>Peugeot</a></li><li><a href='/make/porsche'>Porsche</a></li></ul></div><div class='footer-auto-brend'><ul><li><a href='/make/renault'>Renault</a></li><li><a href='/make/rover'>Rover</a></li><li><a href='/make/saab'>SAAB</a></li><li><a href='/make/ssangyong'>SSangYong</a></li><li><a href='/make/seat'>Seat</a></li><li><a href='/make/smart'>Smart</a></li></ul></div><div class='footer-auto-brend'><ul><li><a href='/make/subaru'>Subaru</a></li><li><a href='/make/suzuki'>Suzuki</a></li><li><a href='/make/tesla'>Tesla</a></li><li><a href='/make/toyota'>Toyota</a></li><li><a href='/make/volkswagen'>Volkswagen</a></li><li><a href='/make/volvo'>Volvo</a></li></ul></div><div class='footer-auto-brend'><ul><li><a href='/make/zastava'>Zastava</a></li><li><a href='/make/škoda'>Škoda</a></li></ul></div>
				</div>
			</div>
		</div>
	</body>
</html>
`
const segment_html = `
<!DOCTYPE html>
<html lang="en">
	<head>
		<meta charset="utf-8">
		<title>
			Audi models - AutoManiac		</title>
		<meta content="width=device-width, initial-scale=1.0" name="viewport">
		<meta content="Audi specifications,Audi consumption,Audi acceleration,Audi prices,Audi buying,Audi selling,Audi opinions" name="keywords">
		<meta content="Audi Models - List of all Audi models and information about each one of them" name="description">
		<meta http-equiv="Cache-Control" content="no-cache, max-age=0, must-revalidate, no-store">
		<meta property="og:site_name" content="AutoManiac"/>
		<meta property="og:image" content="/resources/images/make/cover/audi.jpg" />
		<meta property="og:image:width" content="700"/>
		<meta property="og:image:height" content="300"/>
		<meta property="og:description" content="Audi Models - List of all Audi models and information about each one of them" />
		<meta property="og:title" content="Audi models" />
		<meta property="fb:app_id" content="290604621342699" />
		<meta property="fb:admins" content="1276696339"/>
		<link rel="image_src" type="image/jpeg" href="/resources/images/make/cover/audi.jpg" />
		<link rel="shortcut icon" href="/resources/images/gui/favicon.png" type="image/x-icon">
		<link rel="icon" href="/resources/images/gui/favicon.png" type="image/x-icon">
			<script type="text/javascript" src="/resources/js/jquery-1.8.3.js"></script>
	<script type="text/javascript" src="/resources/js/jquery-ui.js"></script>
	<script type="text/javascript" src="/resources/js/scripts.js?version=v2.1.0"></script>
	<script type="text/javascript" src="/resources/js/jquery.chosen.min.js"></script>
	<script src="https://apis.google.com/js/platform.js" async defer></script>
	<script async src="https://securepubads.g.doubleclick.net/tag/js/gpt.js"></script>
	<!--
	<script async src="//pagead2.googlesyndication.com/pagead/js/adsbygoogle.js"></script>
	<script type="text/javascript" src="/resources/js/adsense.js"></script>
	-->
	<script async defer crossorigin="anonymous" src="https://connect.facebook.net/en_GB/sdk.js#xfbml=1&version=v4.0"></script>
	
						
	<link rel="stylesheet" href="/resources/css/styles.css?version=v2.1.6" type="text/css" media="screen" charset="utf-8" />
	<link rel="stylesheet" href="/resources/css/chosen.css" type="text/css" media="screen" charset="utf-8" />
	<link rel="stylesheet" href="/resources/css/font-awesome.min.css" >
	<link rel="stylesheet" href="https://fonts.googleapis.com/css?family=Open+Sans:300,300i,400,400i,700,700i|Raleway:300,400,500,700,800|Montserrat:300,400,700">
		<link rel="icon" href="/resources/images/gui/favicon.png" >
		<link rel="shortcut icon" href="/resources/images/gui/favicon.png" type="image/x-icon">
		<link rel="apple-touch-icon" href="images/gui/apple-touch-icon.png">
	</head>

	<body root-url="">

		<!-- Load Facebook SDK for JavaScript -->
		<div id="fb-root"></div>
		<script>
		  window.fbAsyncInit = function() {
		    FB.init({
		      appId      : '290604621342699',
		      xfbml      : true,
		      version    : 'v2.8'
		    });
		  };

		  (function(d, s, id){
		     var js, fjs = d.getElementsByTagName(s)[0];
		     if (d.getElementById(id)) {return;}
		     js = d.createElement(s); js.id = id;
		     js.src = "//connect.facebook.net/en_US/sdk.js";
		     fjs.parentNode.insertBefore(js, fjs);
		   }(document, 'script', 'facebook-jssdk'));

		  (function(w,d,u,h,s){
		    h=d.getElementsByTagName('head')[0];
		    s=d.createElement('script');
		    s.async=1;
		    s.src=u+'/sdk.js';
		    h.appendChild(s);
		  })(window,document,'https://aff.carvertical.com');

			
			  (function(i,s,o,g,r,a,m){i['GoogleAnalyticsObject']=r;i[r]=i[r]||function(){
			  (i[r].q=i[r].q||[]).push(arguments)},i[r].l=1*new Date();a=s.createElement(o),
			  m=s.getElementsByTagName(o)[0];a.async=1;a.src=g;m.parentNode.insertBefore(a,m)
			  })(window,document,'script','//www.google-analytics.com/analytics.js','ga');
			
			  ga('create', 'UA-75322623-2', 'auto');
	  		  ga('send', 'pageview');

	  		
		
		</script>

		<div class="menu-wrap">
			<div class="topnav" id="myTopnav">
				<div class="top-bar">
					<div class="top-bar-wrap">
						<div class="top-bar-left">
							<a href="/about" title="about">about</a>
							<a href="/legal" title="legal note">legal note</a>
							<a href="/contact" title="contact us">contact us</a>
							<span style="color: #fff; font-weight: bold;"></span>
						</div>
						<div class="top-bar-right">
							<a href="#" title="">
								Login <i class="fa fa-user" aria-hidden="true"></i>
							</a>
							<a href="https://www.automanie.org" title="">
								deutsch (D, AT, CH)<i class="fa fa-globe" aria-hidden="true"></i>
							</a>
							<a href="https://www.automanijak.com" title="">
								srpski (ex-yugoslavia)<i class="fa fa-flag" aria-hidden="true"></i>
							</a>
						</div>
					</div>
				</div>
				<a href="#" class="icon"><i class="fa fa-bars"></i></a>
				<div class="bottom-bar">
					<div class="bottom-bar-wrap">
						<div class="bottom-bar-left">
							<div class="logo">
								<a href="/" title="">
									<img src="/resources/images/gui/logo.png" alt="" title="">
								</a>
							</div>
						</div>
						<div class="bottom-bar-right">
							<a href="/virtual-adviser">Virtual adviser</a>
							<a href="/specs">Specifications by model</a>
							<a href="/compare">Compare two cars</a>
							<a href="/badges">Automotive badges</a>
							<a href="/blog">Blog</a>
						</div>
					</div>
				</div>
			</div>
		</div>
		<div id="main-wrapper">
			<div class="left-ad">
				<script async src="//pagead2.googlesyndication.com/pagead/js/adsbygoogle.js"></script>
<!-- branding -->
<ins class="adsbygoogle branding-ad"
     style="width:300px;height:600px"
     data-ad-client="ca-pub-2798966384669812"
     data-ad-slot="1676022217"></ins>
<script>
(adsbygoogle = window.adsbygoogle || []).push({});
</script>				<script async src="https://pagead2.googlesyndication.com/pagead/js/adsbygoogle.js"></script>
<!-- massive-header -->
<ins class="adsbygoogle massive-header-ad"
     data-ad-client="ca-pub-2798966384669812"
     data-ad-slot="9878901748"
     data-full-width-responsive="true"></ins>
<script>
     (adsbygoogle = window.adsbygoogle || []).push({});
</script>			</div>
			<div class="right-ad">
				<script async src="//pagead2.googlesyndication.com/pagead/js/adsbygoogle.js"></script>
<!-- branding -->
<ins class="adsbygoogle branding-ad"
     style="width:300px;height:600px"
     data-ad-client="ca-pub-2798966384669812"
     data-ad-slot="1676022217"></ins>
<script>
(adsbygoogle = window.adsbygoogle || []).push({});
</script>			</div>
			<div id="wrapper" class="clearfix">	<div id="sidebar">
<div id="sidebar-dodatno" class="search-box">
	<div class="sidebar-dodatno-wrap">
		<script async src="https://cse.google.com/cse.js?cx=016645478549947938819:d5ilzoxfr9e"></script>
		<script>
			window.onload = function(){
				$(".gsc-input").attr("placeholder", "Search this website...");
			}
		</script>
		<div class="gcse-search"></div>
	</div>
</div>	
	<div class="najpopulariji-wrap-bottom">
		<div class="sidebar-najpopularniji-naslov">
			<h2>Most popular models</h2>
		</div>
		<div class="najpop-item-wrap">
				
	<div class="najpop-item">
		<div class="najpop-item-name">
			<a href="/model/audi/221/audi-a4-2004" title="2004 Audi A4 (B7)">
				2004 Audi A4 (B7)			</a>
		</div>
		<div class="najpop-item-img">
			<a href="/model/audi/221/audi-a4-2004" title="2004 Audi A4 (B7)">
				<img src="/resources/images/variant/41/thumb_a4_5.jpg" 	title="Audi A4"
					 alt="Audi A4"
				 	 class="thumb-image">
			</a>
			<div class="ocena"><span>4.0</span>
			</div>
		</div>
		<div class="najpop-item-val">
			from: 			</br>
			<strong>2.500</strong>
			</br>EUR</div>
	</div>	
	<div class="najpop-item">
		<div class="najpop-item-name">
			<a href="/model/audi/223/audi-a4-2011" title="2011 Audi A4 (B8 restyle)">
				2011 Audi A4 (B8 restyle)			</a>
		</div>
		<div class="najpop-item-img">
			<a href="/model/audi/223/audi-a4-2011" title="2011 Audi A4 (B8 restyle)">
				<img src="/resources/images/variant/46/thumb_a4_10.jpg" 	title="Audi A4"
					 alt="Audi A4"
				 	 class="thumb-image">
			</a>
			<div class="ocena"><span>4.3</span>
			</div>
		</div>
		<div class="najpop-item-val">
			from: 			</br>
			<strong>8.300</strong>
			</br>EUR</div>
	</div>	
	<div class="najpop-item">
		<div class="najpop-item-name">
			<a href="/model/audi/217/audi-a3-2008" title="2008 Audi A3 (8P restyle II)">
				2008 Audi A3 (8P restyle II)			</a>
		</div>
		<div class="najpop-item-img">
			<a href="/model/audi/217/audi-a3-2008" title="2008 Audi A3 (8P restyle II)">
				<img src="/resources/images/variant/31/thumb_a3_7.jpg" 	title="Audi A3"
					 alt="Audi A3"
				 	 class="thumb-image">
			</a>
			<div class="ocena"><span>4.0</span>
			</div>
		</div>
		<div class="najpop-item-val">
			from: 			</br>
			<strong>4.500</strong>
			</br>EUR</div>
	</div>	
	<div class="najpop-item">
		<div class="najpop-item-name">
			<a href="/model/audi/220/audi-a4-2001" title="2001 Audi A4 (B6)">
				2001 Audi A4 (B6)			</a>
		</div>
		<div class="najpop-item-img">
			<a href="/model/audi/220/audi-a4-2001" title="2001 Audi A4 (B6)">
				<img src="/resources/images/variant/39/thumb_a4_4.jpg" 	title="Audi A4"
					 alt="Audi A4"
				 	 class="thumb-image">
			</a>
			<div class="ocena"><span>4.0</span>
			</div>
		</div>
		<div class="najpop-item-val">
			from: 			</br>
			<strong>1.600</strong>
			</br>EUR</div>
	</div>	
	<div class="najpop-item">
		<div class="najpop-item-name">
			<a href="/model/audi/216/audi-a3-2005" title="2005 Audi A3 (8P restyle I)">
				2005 Audi A3 (8P restyle I)			</a>
		</div>
		<div class="najpop-item-img">
			<a href="/model/audi/216/audi-a3-2005" title="2005 Audi A3 (8P restyle I)">
				<img src="/resources/images/variant/29/thumb_a3_5.jpg" 	title="Audi A3"
					 alt="Audi A3"
				 	 class="thumb-image">
			</a>
			<div class="ocena"><span>4.4</span>
			</div>
		</div>
		<div class="najpop-item-val">
			from: 			</br>
			<strong>2.800</strong>
			</br>EUR</div>
	</div>		</div>
	</div>

	<div class="side-ad">
	<script async src="//pagead2.googlesyndication.com/pagead/js/adsbygoogle.js"></script>
	<ins class="adsbygoogle side-ad-vertical"
	     style="width:160px;height:600px"
	     data-ad-client="ca-pub-2798966384669812"
	     data-ad-slot="9565981886"></ins>
	<script>
	     (adsbygoogle = window.adsbygoogle || []).push({});
	</script>

	<ins class="adsbygoogle side-ad-horizontal"
	     data-ad-client="ca-pub-2798966384669812"
	     data-ad-slot="8347642282"
	     data-full-width-responsive="true"></ins>
	<script>
	     (adsbygoogle = window.adsbygoogle || []).push({});
	</script>
</div>
	</div>
	<div id="page-wrap">
<div id="breadcrumb-wrap">
	<div class="breadcrumb-nav"><a href='/'>HOME page</a> / <a href='/specs'>specifications</a> / Audi models</div>
	<div id="sidebar-dodatno" class="search-box">
	<div class="sidebar-dodatno-wrap">
		<script async src="https://cse.google.com/cse.js?cx=016645478549947938819:d5ilzoxfr9e"></script>
		<script>
			window.onload = function(){
				$(".gsc-input").attr("placeholder", "Search this website...");
			}
		</script>
		<div class="gcse-search"></div>
	</div>
</div>	<div class="facebook">
				<div class="fb-like" data-href="https://www.automaniac.org/make/audi/segment" data-width="" data-layout="button_count" data-action="like" data-size="small" data-show-faces="true" data-share="true"></div>
	</div>
</div>	<div id="modeli-auto-lista">
		<div class="box-wrap">
			<div class="model-logo">
				<img src="/resources/images/make/audi.jpg"
					 alt="Audi models" />
			</div>
			<div class="model-naslov">
				Audi models				<span>List of all models from Audi</span>
			</div>
			<div class="model-zemlja">
				<img src="/resources/images/country/DE.png" alt="DE">
			</div>
			<div class="model-slika">
				<img src="/resources/images/make/cover/audi.jpg"
					 alt="Audi models" />
			</div>
		</div>
				<div>
			<h2><i class="fa fa-calendar" aria-hidden="true"></i>View by: <strong>segment</strong></h2>
		</div>
		<div class="box-wrap">
			<div class="segmenti">
				<div class="segment selected">
					<a href="/make/audi/segment"
					   title="View by<br/>segment"
					   alt="View by<br/>segment">
					<i class="fa fa-pie-chart" aria-hidden="true"></i>
						<p>View by<br/>segment</p>
					</a>
				</div>
				<div class="segment ">
					<a href="/make/audi/year"
					   title="View by<br/>year"
					   alt="View by<br/>year">
					<i class="fa fa-area-chart" aria-hidden="true"></i>
						<p>View by<br/>year</p>
					</a>
				</div>
				<div class="segment ">
					<a href="/make/audi/year-segment"
					   title="View by<br/>segment and year"
					   alt="View by<br/>segment and year">
						<i class="fa fa-bar-chart" aria-hidden="true"></i>
						<p>View by<br/>segment and year</p>
					</a>
				</div>
			</div>
		</div>


<div class="box-wrap">

	<div class="proizvodnja-wrap">

</div>
<div class="proizvodnja-wrap">


		<div class="proizvodnja-box variant-header accordion-controller segment-header"
			 related-accordion="accordion-B"
			 segment="B"
		 >
			<div class="proizvodnja-godina">
				B - City car			</div>
			<div class="proizvodnja-model">
				2 models			</div>
		</div>
		<div id="accordion-B" class="variant-accordion ">
							<div class="auto-wrap">
					<div class="auto-slika">
						<div class="auto-score-wrap clearfix">
														<div class="score-left">
								<div class="score4b">
									<p>drivers'</p>
									<span>3.6</span>
									<p>rating</p>
								</div>
							</div>
							<div class="score-right clearfix">
								<div class="score5a"></div>
								<div class="score4a"></div>
								<div class="score3a"></div>
								<div class="score2a"></div>
								<div class="score1a"></div>
							</div>
						</div>
						<a href="/model/audi/212/audi-a1-2010" title="">
							<img src="/resources/images/variant/22/thumb_a1_2.jpg"
								 alt=""
								 class="thumb-image"/>
						</a>
					</div>
					<div class="auto-info-wrap">
						<div class="auto-naziv">
							<a href="/model/audi/212/audi-a1-2010" title="2010 Audi A1 (8X)">
								2010 Audi A1 (8X)							</a>
						</div>
						<div class="stars">
							<div class="stars5"></div>
						</div>
						<div class="euro-ncap">
							EuroNCAP:&nbsp;
							90% occupant safety						</div>
						<div>
							<b>
								<a href="/model/audi/212/audi-a1-2010" title="2010 Audi A1 (8X)" class="orange">
							   	   	more details...						   	   </a>
					   	   </b>
						</div>
					</div>
				</div>
							<div class="auto-wrap">
					<div class="auto-slika">
						<div class="auto-score-wrap clearfix">
														<div class="score-left">
								<div class="score3b">
									<p>drivers'</p>
									<span>3.0</span>
									<p>rating</p>
								</div>
							</div>
							<div class="score-right clearfix">
								<div class="score5a"></div>
								<div class="score4a"></div>
								<div class="score3a"></div>
								<div class="score2a"></div>
								<div class="score1a"></div>
							</div>
						</div>
						<a href="/model/audi/1235/audi-a1-2014" title="">
							<img src="/resources/images/variant/1853/thumb_a1_3.jpg"
								 alt=""
								 class="thumb-image"/>
						</a>
					</div>
					<div class="auto-info-wrap">
						<div class="auto-naziv">
							<a href="/model/audi/1235/audi-a1-2014" title="2014 Audi A1 (8X restyle)">
								2014 Audi A1 (8X restyle)							</a>
						</div>
						<div class="stars">
							<div class="stars5"></div>
						</div>
						<div class="euro-ncap">
							EuroNCAP:&nbsp;
							90% occupant safety						</div>
						<div>
							<b>
								<a href="/model/audi/1235/audi-a1-2014" title="2014 Audi A1 (8X restyle)" class="orange">
							   	   	more details...						   	   </a>
					   	   </b>
						</div>
					</div>
				</div>
					</div>
</div>
<div class="proizvodnja-wrap">


		<div class="proizvodnja-box variant-header accordion-controller segment-header"
			 related-accordion="accordion-C"
			 segment="C"
		 >
			<div class="proizvodnja-godina">
				C - Small family car			</div>
			<div class="proizvodnja-model">
				7 models			</div>
		</div>
		<div id="accordion-C" class="variant-accordion ">
							<div class="auto-wrap">
					<div class="auto-slika">
						<div class="auto-score-wrap clearfix">
														<div class="score-left">
								<div class="score3b">
									<p>drivers'</p>
									<span>2.6</span>
									<p>rating</p>
								</div>
							</div>
							<div class="score-right clearfix">
								<div class="score5a"></div>
								<div class="score4a"></div>
								<div class="score3a"></div>
								<div class="score2a"></div>
								<div class="score1a"></div>
							</div>
						</div>
						<a href="/model/audi/214/audi-a3-1999" title="">
							<img src="/resources/images/variant/25/thumb_a3_1.jpg"
								 alt=""
								 class="thumb-image"/>
						</a>
					</div>
					<div class="auto-info-wrap">
						<div class="auto-naziv">
							<a href="/model/audi/214/audi-a3-1999" title="1999 Audi A3 (8L restyle)">
								1999 Audi A3 (8L restyle)							</a>
						</div>
						<div class="stars">
							<div class="stars4"></div>
						</div>
						<div class="euro-ncap">
							EuroNCAP:&nbsp;
							~70% occupant safety						</div>
						<div>
							<b>
								<a href="/model/audi/214/audi-a3-1999" title="1999 Audi A3 (8L restyle)" class="orange">
							   	   	more details...						   	   </a>
					   	   </b>
						</div>
					</div>
				</div>
							<div class="auto-wrap">
					<div class="auto-slika">
						<div class="auto-score-wrap clearfix">
														<div class="score-left">
								<div class="score4b">
									<p>drivers'</p>
									<span>4.1</span>
									<p>rating</p>
								</div>
							</div>
							<div class="score-right clearfix">
								<div class="score5a"></div>
								<div class="score4a"></div>
								<div class="score3a"></div>
								<div class="score2a"></div>
								<div class="score1a"></div>
							</div>
						</div>
						<a href="/model/audi/215/audi-a3-2003" title="">
							<img src="/resources/images/variant/27/thumb_a3_3.jpg"
								 alt=""
								 class="thumb-image"/>
						</a>
					</div>
					<div class="auto-info-wrap">
						<div class="auto-naziv">
							<a href="/model/audi/215/audi-a3-2003" title="2003 Audi A3 (8P)">
								2003 Audi A3 (8P)							</a>
						</div>
						<div class="stars">
							<div class="stars4"></div>
						</div>
						<div class="euro-ncap">
							EuroNCAP:&nbsp;
							~70% occupant safety						</div>
						<div>
							<b>
								<a href="/model/audi/215/audi-a3-2003" title="2003 Audi A3 (8P)" class="orange">
							   	   	more details...						   	   </a>
					   	   </b>
						</div>
					</div>
				</div>
							<div class="auto-wrap">
					<div class="auto-slika">
						<div class="auto-score-wrap clearfix">
														<div class="score-left">
								<div class="score4b">
									<p>drivers'</p>
									<span>4.4</span>
									<p>rating</p>
								</div>
							</div>
							<div class="score-right clearfix">
								<div class="score5a"></div>
								<div class="score4a"></div>
								<div class="score3a"></div>
								<div class="score2a"></div>
								<div class="score1a"></div>
							</div>
						</div>
						<a href="/model/audi/216/audi-a3-2005" title="">
							<img src="/resources/images/variant/29/thumb_a3_5.jpg"
								 alt=""
								 class="thumb-image"/>
						</a>
					</div>
					<div class="auto-info-wrap">
						<div class="auto-naziv">
							<a href="/model/audi/216/audi-a3-2005" title="2005 Audi A3 (8P restyle I)">
								2005 Audi A3 (8P restyle I)							</a>
						</div>
						<div class="stars">
							<div class="stars4"></div>
						</div>
						<div class="euro-ncap">
							EuroNCAP:&nbsp;
							~70% occupant safety						</div>
						<div>
							<b>
								<a href="/model/audi/216/audi-a3-2005" title="2005 Audi A3 (8P restyle I)" class="orange">
							   	   	more details...						   	   </a>
					   	   </b>
						</div>
					</div>
				</div>
							<div class="auto-wrap">
					<div class="auto-slika">
						<div class="auto-score-wrap clearfix">
														<div class="score-left">
								<div class="score4b">
									<p>drivers'</p>
									<span>4.0</span>
									<p>rating</p>
								</div>
							</div>
							<div class="score-right clearfix">
								<div class="score5a"></div>
								<div class="score4a"></div>
								<div class="score3a"></div>
								<div class="score2a"></div>
								<div class="score1a"></div>
							</div>
						</div>
						<a href="/model/audi/217/audi-a3-2008" title="">
							<img src="/resources/images/variant/31/thumb_a3_7.jpg"
								 alt=""
								 class="thumb-image"/>
						</a>
					</div>
					<div class="auto-info-wrap">
						<div class="auto-naziv">
							<a href="/model/audi/217/audi-a3-2008" title="2008 Audi A3 (8P restyle II)">
								2008 Audi A3 (8P restyle II)							</a>
						</div>
						<div class="stars">
							<div class="stars4"></div>
						</div>
						<div class="euro-ncap">
							EuroNCAP:&nbsp;
							~70% occupant safety						</div>
						<div>
							<b>
								<a href="/model/audi/217/audi-a3-2008" title="2008 Audi A3 (8P restyle II)" class="orange">
							   	   	more details...						   	   </a>
					   	   </b>
						</div>
					</div>
				</div>
							<div class="auto-wrap">
					<div class="auto-slika">
						<div class="auto-score-wrap clearfix">
														<div class="score-left">
								<div class="score5b">
									<p>drivers'</p>
									<span>4.6</span>
									<p>rating</p>
								</div>
							</div>
							<div class="score-right clearfix">
								<div class="score5a"></div>
								<div class="score4a"></div>
								<div class="score3a"></div>
								<div class="score2a"></div>
								<div class="score1a"></div>
							</div>
						</div>
						<a href="/model/audi/218/audi-a3-2012" title="">
							<img src="/resources/images/variant/34/thumb_a3_10.jpg"
								 alt=""
								 class="thumb-image"/>
						</a>
					</div>
					<div class="auto-info-wrap">
						<div class="auto-naziv">
							<a href="/model/audi/218/audi-a3-2012" title="2012 Audi A3 (8V)">
								2012 Audi A3 (8V)							</a>
						</div>
						<div class="stars">
							<div class="stars5"></div>
						</div>
						<div class="euro-ncap">
							EuroNCAP:&nbsp;
							95% occupant safety						</div>
						<div>
							<b>
								<a href="/model/audi/218/audi-a3-2012" title="2012 Audi A3 (8V)" class="orange">
							   	   	more details...						   	   </a>
					   	   </b>
						</div>
					</div>
				</div>
							<div class="auto-wrap">
					<div class="auto-slika">
						<div class="auto-score-wrap clearfix">
														<div class="score-left">
								<div class="score4b">
									<p>drivers'</p>
									<span>4.5</span>
									<p>rating</p>
								</div>
							</div>
							<div class="score-right clearfix">
								<div class="score5a"></div>
								<div class="score4a"></div>
								<div class="score3a"></div>
								<div class="score2a"></div>
								<div class="score1a"></div>
							</div>
						</div>
						<a href="/model/audi/1239/audi-a3-2016" title="">
							<img src="/resources/images/variant/1861/thumb_a3_13.jpg"
								 alt=""
								 class="thumb-image"/>
						</a>
					</div>
					<div class="auto-info-wrap">
						<div class="auto-naziv">
							<a href="/model/audi/1239/audi-a3-2016" title="2016 Audi A3 (8V restyle)">
								2016 Audi A3 (8V restyle)							</a>
						</div>
						<div class="stars">
							<div class="stars5"></div>
						</div>
						<div class="euro-ncap">
							EuroNCAP:&nbsp;
							95% occupant safety						</div>
						<div>
							<b>
								<a href="/model/audi/1239/audi-a3-2016" title="2016 Audi A3 (8V restyle)" class="orange">
							   	   	more details...						   	   </a>
					   	   </b>
						</div>
					</div>
				</div>
							<div class="auto-wrap">
					<div class="auto-slika">
						<div class="auto-score-wrap clearfix">
														<div class="score-left">
								<div class="score3b">
									<p>drivers'</p>
									<span>3.0</span>
									<p>rating</p>
								</div>
							</div>
							<div class="score-right clearfix">
								<div class="score5a"></div>
								<div class="score4a"></div>
								<div class="score3a"></div>
								<div class="score2a"></div>
								<div class="score1a"></div>
							</div>
						</div>
						<a href="/model/audi/1459/audi-a3-2020" title="">
							<img src="/resources/images/variant/2213/thumb_2020_audi_a3_sportback.jpg"
								 alt=""
								 class="thumb-image"/>
						</a>
					</div>
					<div class="auto-info-wrap">
						<div class="auto-naziv">
							<a href="/model/audi/1459/audi-a3-2020" title="2020 Audi A3 (8Y)">
								2020 Audi A3 (8Y)							</a>
						</div>
						<div class="stars">
							<div class="stars"></div>
						</div>
						<div class="euro-ncap">
							EuroNCAP:&nbsp;
							not tested						</div>
						<div>
							<b>
								<a href="/model/audi/1459/audi-a3-2020" title="2020 Audi A3 (8Y)" class="orange">
							   	   	more details...						   	   </a>
					   	   </b>
						</div>
					</div>
				</div>
					</div>
</div>
<div class="proizvodnja-wrap">


		<div class="proizvodnja-box variant-header accordion-controller segment-header"
			 related-accordion="accordion-D"
			 segment="D"
		 >
			<div class="proizvodnja-godina">
				D - Large family car			</div>
			<div class="proizvodnja-model">
				10 models			</div>
		</div>
		<div id="accordion-D" class="variant-accordion ">
							<div class="auto-wrap">
					<div class="auto-slika">
						<div class="auto-score-wrap clearfix">
														<div class="score-left">
								<div class="score4b">
									<p>drivers'</p>
									<span>4.2</span>
									<p>rating</p>
								</div>
							</div>
							<div class="score-right clearfix">
								<div class="score5a"></div>
								<div class="score4a"></div>
								<div class="score3a"></div>
								<div class="score2a"></div>
								<div class="score1a"></div>
							</div>
						</div>
						<a href="/model/audi/219/audi-a4-1999" title="">
							<img src="/resources/images/variant/37/thumb_a4_1.jpg"
								 alt=""
								 class="thumb-image"/>
						</a>
					</div>
					<div class="auto-info-wrap">
						<div class="auto-naziv">
							<a href="/model/audi/219/audi-a4-1999" title="1999 Audi A4 (B5 restyle)">
								1999 Audi A4 (B5 restyle)							</a>
						</div>
						<div class="stars">
							<div class="stars3"></div>
						</div>
						<div class="euro-ncap">
							EuroNCAP:&nbsp;
							~50% occupant safety						</div>
						<div>
							<b>
								<a href="/model/audi/219/audi-a4-1999" title="1999 Audi A4 (B5 restyle)" class="orange">
							   	   	more details...						   	   </a>
					   	   </b>
						</div>
					</div>
				</div>
							<div class="auto-wrap">
					<div class="auto-slika">
						<div class="auto-score-wrap clearfix">
														<div class="score-left">
								<div class="score4b">
									<p>drivers'</p>
									<span>4.0</span>
									<p>rating</p>
								</div>
							</div>
							<div class="score-right clearfix">
								<div class="score5a"></div>
								<div class="score4a"></div>
								<div class="score3a"></div>
								<div class="score2a"></div>
								<div class="score1a"></div>
							</div>
						</div>
						<a href="/model/audi/220/audi-a4-2001" title="">
							<img src="/resources/images/variant/39/thumb_a4_4.jpg"
								 alt=""
								 class="thumb-image"/>
						</a>
					</div>
					<div class="auto-info-wrap">
						<div class="auto-naziv">
							<a href="/model/audi/220/audi-a4-2001" title="2001 Audi A4 (B6)">
								2001 Audi A4 (B6)							</a>
						</div>
						<div class="stars">
							<div class="stars4"></div>
						</div>
						<div class="euro-ncap">
							EuroNCAP:&nbsp;
							~70% occupant safety						</div>
						<div>
							<b>
								<a href="/model/audi/220/audi-a4-2001" title="2001 Audi A4 (B6)" class="orange">
							   	   	more details...						   	   </a>
					   	   </b>
						</div>
					</div>
				</div>
							<div class="auto-wrap">
					<div class="auto-slika">
						<div class="auto-score-wrap clearfix">
														<div class="score-left">
								<div class="score4b">
									<p>drivers'</p>
									<span>4.0</span>
									<p>rating</p>
								</div>
							</div>
							<div class="score-right clearfix">
								<div class="score5a"></div>
								<div class="score4a"></div>
								<div class="score3a"></div>
								<div class="score2a"></div>
								<div class="score1a"></div>
							</div>
						</div>
						<a href="/model/audi/221/audi-a4-2004" title="">
							<img src="/resources/images/variant/41/thumb_a4_5.jpg"
								 alt=""
								 class="thumb-image"/>
						</a>
					</div>
					<div class="auto-info-wrap">
						<div class="auto-naziv">
							<a href="/model/audi/221/audi-a4-2004" title="2004 Audi A4 (B7)">
								2004 Audi A4 (B7)							</a>
						</div>
						<div class="stars">
							<div class="stars4"></div>
						</div>
						<div class="euro-ncap">
							EuroNCAP:&nbsp;
							~70% occupant safety						</div>
						<div>
							<b>
								<a href="/model/audi/221/audi-a4-2004" title="2004 Audi A4 (B7)" class="orange">
							   	   	more details...						   	   </a>
					   	   </b>
						</div>
					</div>
				</div>
							<div class="auto-wrap">
					<div class="auto-slika">
						<div class="auto-score-wrap clearfix">
														<div class="score-left">
								<div class="score4b">
									<p>drivers'</p>
									<span>4.0</span>
									<p>rating</p>
								</div>
							</div>
							<div class="score-right clearfix">
								<div class="score5a"></div>
								<div class="score4a"></div>
								<div class="score3a"></div>
								<div class="score2a"></div>
								<div class="score1a"></div>
							</div>
						</div>
						<a href="/model/audi/222/audi-a4-2007" title="">
							<img src="/resources/images/variant/44/thumb_a4_8.jpg"
								 alt=""
								 class="thumb-image"/>
						</a>
					</div>
					<div class="auto-info-wrap">
						<div class="auto-naziv">
							<a href="/model/audi/222/audi-a4-2007" title="2007 Audi A4 (B8)">
								2007 Audi A4 (B8)							</a>
						</div>
						<div class="stars">
							<div class="stars5"></div>
						</div>
						<div class="euro-ncap">
							EuroNCAP:&nbsp;
							~90% occupant safety						</div>
						<div>
							<b>
								<a href="/model/audi/222/audi-a4-2007" title="2007 Audi A4 (B8)" class="orange">
							   	   	more details...						   	   </a>
					   	   </b>
						</div>
					</div>
				</div>
							<div class="auto-wrap">
					<div class="auto-slika">
						<div class="auto-score-wrap clearfix">
														<div class="score-left">
								<div class="score4b">
									<p>drivers'</p>
									<span>4.3</span>
									<p>rating</p>
								</div>
							</div>
							<div class="score-right clearfix">
								<div class="score5a"></div>
								<div class="score4a"></div>
								<div class="score3a"></div>
								<div class="score2a"></div>
								<div class="score1a"></div>
							</div>
						</div>
						<a href="/model/audi/223/audi-a4-2011" title="">
							<img src="/resources/images/variant/46/thumb_a4_10.jpg"
								 alt=""
								 class="thumb-image"/>
						</a>
					</div>
					<div class="auto-info-wrap">
						<div class="auto-naziv">
							<a href="/model/audi/223/audi-a4-2011" title="2011 Audi A4 (B8 restyle)">
								2011 Audi A4 (B8 restyle)							</a>
						</div>
						<div class="stars">
							<div class="stars5"></div>
						</div>
						<div class="euro-ncap">
							EuroNCAP:&nbsp;
							93% occupant safety						</div>
						<div>
							<b>
								<a href="/model/audi/223/audi-a4-2011" title="2011 Audi A4 (B8 restyle)" class="orange">
							   	   	more details...						   	   </a>
					   	   </b>
						</div>
					</div>
				</div>
							<div class="auto-wrap">
					<div class="auto-slika">
						<div class="auto-score-wrap clearfix">
														<div class="score-left">
								<div class="score3b">
									<p>drivers'</p>
									<span>2.8</span>
									<p>rating</p>
								</div>
							</div>
							<div class="score-right clearfix">
								<div class="score5a"></div>
								<div class="score4a"></div>
								<div class="score3a"></div>
								<div class="score2a"></div>
								<div class="score1a"></div>
							</div>
						</div>
						<a href="/model/audi/1187/audi-a4-2015" title="">
							<img src="/resources/images/variant/1780/thumb_a4_25.jpg"
								 alt=""
								 class="thumb-image"/>
						</a>
					</div>
					<div class="auto-info-wrap">
						<div class="auto-naziv">
							<a href="/model/audi/1187/audi-a4-2015" title="2015 Audi A4 (B9)">
								2015 Audi A4 (B9)							</a>
						</div>
						<div class="stars">
							<div class="stars5"></div>
						</div>
						<div class="euro-ncap">
							EuroNCAP:&nbsp;
							90% occupant safety						</div>
						<div>
							<b>
								<a href="/model/audi/1187/audi-a4-2015" title="2015 Audi A4 (B9)" class="orange">
							   	   	more details...						   	   </a>
					   	   </b>
						</div>
					</div>
				</div>
							<div class="auto-wrap">
					<div class="auto-slika">
						<div class="auto-score-wrap clearfix">
														<div class="score-left">
								<div class="score3b">
									<p>drivers'</p>
									<span>3.0</span>
									<p>rating</p>
								</div>
							</div>
							<div class="score-right clearfix">
								<div class="score5a"></div>
								<div class="score4a"></div>
								<div class="score3a"></div>
								<div class="score2a"></div>
								<div class="score1a"></div>
							</div>
						</div>
						<a href="/model/audi/1394/audi-a4-2018" title="">
							<img src="/resources/images/variant/2107/thumb_2018_audi_a4.jpg"
								 alt=""
								 class="thumb-image"/>
						</a>
					</div>
					<div class="auto-info-wrap">
						<div class="auto-naziv">
							<a href="/model/audi/1394/audi-a4-2018" title="2018 Audi A4 (B9 restyle)">
								2018 Audi A4 (B9 restyle)							</a>
						</div>
						<div class="stars">
							<div class="stars5"></div>
						</div>
						<div class="euro-ncap">
							EuroNCAP:&nbsp;
							90% occupant safety						</div>
						<div>
							<b>
								<a href="/model/audi/1394/audi-a4-2018" title="2018 Audi A4 (B9 restyle)" class="orange">
							   	   	more details...						   	   </a>
					   	   </b>
						</div>
					</div>
				</div>
							<div class="auto-wrap">
					<div class="auto-slika">
						<div class="auto-score-wrap clearfix">
														<div class="score-left">
								<div class="score1b">
									<p>drivers'</p>
									<span>1.0</span>
									<p>rating</p>
								</div>
							</div>
							<div class="score-right clearfix">
								<div class="score5a"></div>
								<div class="score4a"></div>
								<div class="score3a"></div>
								<div class="score2a"></div>
								<div class="score1a"></div>
							</div>
						</div>
						<a href="/model/audi/226/audi-a5-sportback-2009" title="">
							<img src="/resources/images/variant/53/thumb_a5_5.jpg"
								 alt=""
								 class="thumb-image"/>
						</a>
					</div>
					<div class="auto-info-wrap">
						<div class="auto-naziv">
							<a href="/model/audi/226/audi-a5-sportback-2009" title="2009 Audi A5 Sportback (8T)">
								2009 Audi A5 Sportback (8T)							</a>
						</div>
						<div class="stars">
							<div class="stars"></div>
						</div>
						<div class="euro-ncap">
							EuroNCAP:&nbsp;
							not tested						</div>
						<div>
							<b>
								<a href="/model/audi/226/audi-a5-sportback-2009" title="2009 Audi A5 Sportback (8T)" class="orange">
							   	   	more details...						   	   </a>
					   	   </b>
						</div>
					</div>
				</div>
							<div class="auto-wrap">
					<div class="auto-slika">
						<div class="auto-score-wrap clearfix">
														<div class="score-left">
								<div class="score4b">
									<p>drivers'</p>
									<span>4.1</span>
									<p>rating</p>
								</div>
							</div>
							<div class="score-right clearfix">
								<div class="score5a"></div>
								<div class="score4a"></div>
								<div class="score3a"></div>
								<div class="score2a"></div>
								<div class="score1a"></div>
							</div>
						</div>
						<a href="/model/audi/227/audi-a5-sportback-2011" title="">
							<img src="/resources/images/variant/54/thumb_a5_6.jpg"
								 alt=""
								 class="thumb-image"/>
						</a>
					</div>
					<div class="auto-info-wrap">
						<div class="auto-naziv">
							<a href="/model/audi/227/audi-a5-sportback-2011" title="2011 Audi A5 Sportback (8T restyle)">
								2011 Audi A5 Sportback (8T restyle)							</a>
						</div>
						<div class="stars">
							<div class="stars"></div>
						</div>
						<div class="euro-ncap">
							EuroNCAP:&nbsp;
							not tested						</div>
						<div>
							<b>
								<a href="/model/audi/227/audi-a5-sportback-2011" title="2011 Audi A5 Sportback (8T restyle)" class="orange">
							   	   	more details...						   	   </a>
					   	   </b>
						</div>
					</div>
				</div>
							<div class="auto-wrap">
					<div class="auto-slika">
						<div class="auto-score-wrap clearfix">
														<div class="score-left">
								<div class="score3b">
									<p>drivers'</p>
									<span>3.0</span>
									<p>rating</p>
								</div>
							</div>
							<div class="score-right clearfix">
								<div class="score5a"></div>
								<div class="score4a"></div>
								<div class="score3a"></div>
								<div class="score2a"></div>
								<div class="score1a"></div>
							</div>
						</div>
						<a href="/model/audi/1245/audi-a5-sportback-2016" title="">
							<img src="/resources/images/variant/1877/thumb_a5_7.jpg"
								 alt=""
								 class="thumb-image"/>
						</a>
					</div>
					<div class="auto-info-wrap">
						<div class="auto-naziv">
							<a href="/model/audi/1245/audi-a5-sportback-2016" title="2016 Audi A5 Sportback (F5)">
								2016 Audi A5 Sportback (F5)							</a>
						</div>
						<div class="stars">
							<div class="stars5"></div>
						</div>
						<div class="euro-ncap">
							EuroNCAP:&nbsp;
							89% occupant safety						</div>
						<div>
							<b>
								<a href="/model/audi/1245/audi-a5-sportback-2016" title="2016 Audi A5 Sportback (F5)" class="orange">
							   	   	more details...						   	   </a>
					   	   </b>
						</div>
					</div>
				</div>
					</div>
</div>
<div class="proizvodnja-wrap">


		<div class="proizvodnja-box variant-header accordion-controller segment-header"
			 related-accordion="accordion-E"
			 segment="E"
		 >
			<div class="proizvodnja-godina">
				E - Luxury car			</div>
			<div class="proizvodnja-model">
				10 models			</div>
		</div>
		<div id="accordion-E" class="variant-accordion ">
							<div class="auto-wrap">
					<div class="auto-slika">
						<div class="auto-score-wrap clearfix">
														<div class="score-left">
								<div class="score4b">
									<p>drivers'</p>
									<span>4.1</span>
									<p>rating</p>
								</div>
							</div>
							<div class="score-right clearfix">
								<div class="score5a"></div>
								<div class="score4a"></div>
								<div class="score3a"></div>
								<div class="score2a"></div>
								<div class="score1a"></div>
							</div>
						</div>
						<a href="/model/audi/228/audi-a6-1997" title="">
							<img src="/resources/images/variant/55/thumb_a6_1.jpg"
								 alt=""
								 class="thumb-image"/>
						</a>
					</div>
					<div class="auto-info-wrap">
						<div class="auto-naziv">
							<a href="/model/audi/228/audi-a6-1997" title="1997 Audi A6 (C5)">
								1997 Audi A6 (C5)							</a>
						</div>
						<div class="stars">
							<div class="stars3"></div>
						</div>
						<div class="euro-ncap">
							EuroNCAP:&nbsp;
							~50% occupant safety						</div>
						<div>
							<b>
								<a href="/model/audi/228/audi-a6-1997" title="1997 Audi A6 (C5)" class="orange">
							   	   	more details...						   	   </a>
					   	   </b>
						</div>
					</div>
				</div>
							<div class="auto-wrap">
					<div class="auto-slika">
						<div class="auto-score-wrap clearfix">
														<div class="score-left">
								<div class="score4b">
									<p>drivers'</p>
									<span>4.4</span>
									<p>rating</p>
								</div>
							</div>
							<div class="score-right clearfix">
								<div class="score5a"></div>
								<div class="score4a"></div>
								<div class="score3a"></div>
								<div class="score2a"></div>
								<div class="score1a"></div>
							</div>
						</div>
						<a href="/model/audi/229/audi-a6-2001" title="">
							<img src="/resources/images/variant/57/thumb_a6_3.jpg"
								 alt=""
								 class="thumb-image"/>
						</a>
					</div>
					<div class="auto-info-wrap">
						<div class="auto-naziv">
							<a href="/model/audi/229/audi-a6-2001" title="2001 Audi A6 (C5 restyle)">
								2001 Audi A6 (C5 restyle)							</a>
						</div>
						<div class="stars">
							<div class="stars3"></div>
						</div>
						<div class="euro-ncap">
							EuroNCAP:&nbsp;
							~50% occupant safety						</div>
						<div>
							<b>
								<a href="/model/audi/229/audi-a6-2001" title="2001 Audi A6 (C5 restyle)" class="orange">
							   	   	more details...						   	   </a>
					   	   </b>
						</div>
					</div>
				</div>
							<div class="auto-wrap">
					<div class="auto-slika">
						<div class="auto-score-wrap clearfix">
														<div class="score-left">
								<div class="score4b">
									<p>drivers'</p>
									<span>4.2</span>
									<p>rating</p>
								</div>
							</div>
							<div class="score-right clearfix">
								<div class="score5a"></div>
								<div class="score4a"></div>
								<div class="score3a"></div>
								<div class="score2a"></div>
								<div class="score1a"></div>
							</div>
						</div>
						<a href="/model/audi/230/audi-a6-2004" title="">
							<img src="/resources/images/variant/59/thumb_a6_5.jpg"
								 alt=""
								 class="thumb-image"/>
						</a>
					</div>
					<div class="auto-info-wrap">
						<div class="auto-naziv">
							<a href="/model/audi/230/audi-a6-2004" title="2004 Audi A6 (C6)">
								2004 Audi A6 (C6)							</a>
						</div>
						<div class="stars">
							<div class="stars5"></div>
						</div>
						<div class="euro-ncap">
							EuroNCAP:&nbsp;
							~90% occupant safety						</div>
						<div>
							<b>
								<a href="/model/audi/230/audi-a6-2004" title="2004 Audi A6 (C6)" class="orange">
							   	   	more details...						   	   </a>
					   	   </b>
						</div>
					</div>
				</div>
							<div class="auto-wrap">
					<div class="auto-slika">
						<div class="auto-score-wrap clearfix">
														<div class="score-left">
								<div class="score4b">
									<p>drivers'</p>
									<span>4.4</span>
									<p>rating</p>
								</div>
							</div>
							<div class="score-right clearfix">
								<div class="score5a"></div>
								<div class="score4a"></div>
								<div class="score3a"></div>
								<div class="score2a"></div>
								<div class="score1a"></div>
							</div>
						</div>
						<a href="/model/audi/231/audi-a6-2008" title="">
							<img src="/resources/images/variant/61/thumb_a6_7.jpg"
								 alt=""
								 class="thumb-image"/>
						</a>
					</div>
					<div class="auto-info-wrap">
						<div class="auto-naziv">
							<a href="/model/audi/231/audi-a6-2008" title="2008 Audi A6 (C6 restyle)">
								2008 Audi A6 (C6 restyle)							</a>
						</div>
						<div class="stars">
							<div class="stars5"></div>
						</div>
						<div class="euro-ncap">
							EuroNCAP:&nbsp;
							~90% occupant safety						</div>
						<div>
							<b>
								<a href="/model/audi/231/audi-a6-2008" title="2008 Audi A6 (C6 restyle)" class="orange">
							   	   	more details...						   	   </a>
					   	   </b>
						</div>
					</div>
				</div>
							<div class="auto-wrap">
					<div class="auto-slika">
						<div class="auto-score-wrap clearfix">
														<div class="score-left">
								<div class="score4b">
									<p>drivers'</p>
									<span>4.1</span>
									<p>rating</p>
								</div>
							</div>
							<div class="score-right clearfix">
								<div class="score5a"></div>
								<div class="score4a"></div>
								<div class="score3a"></div>
								<div class="score2a"></div>
								<div class="score1a"></div>
							</div>
						</div>
						<a href="/model/audi/232/audi-a6-2011" title="">
							<img src="/resources/images/variant/63/thumb_a6_9.jpg"
								 alt=""
								 class="thumb-image"/>
						</a>
					</div>
					<div class="auto-info-wrap">
						<div class="auto-naziv">
							<a href="/model/audi/232/audi-a6-2011" title="2011 Audi A6 (C7)">
								2011 Audi A6 (C7)							</a>
						</div>
						<div class="stars">
							<div class="stars5"></div>
						</div>
						<div class="euro-ncap">
							EuroNCAP:&nbsp;
							91% occupant safety						</div>
						<div>
							<b>
								<a href="/model/audi/232/audi-a6-2011" title="2011 Audi A6 (C7)" class="orange">
							   	   	more details...						   	   </a>
					   	   </b>
						</div>
					</div>
				</div>
							<div class="auto-wrap">
					<div class="auto-slika">
						<div class="auto-score-wrap clearfix">
														<div class="score-left">
								<div class="score5b">
									<p>drivers'</p>
									<span>4.8</span>
									<p>rating</p>
								</div>
							</div>
							<div class="score-right clearfix">
								<div class="score5a"></div>
								<div class="score4a"></div>
								<div class="score3a"></div>
								<div class="score2a"></div>
								<div class="score1a"></div>
							</div>
						</div>
						<a href="/model/audi/1213/audi-a6-2014" title="">
							<img src="/resources/images/variant/1816/thumb_a6_15.jpg"
								 alt=""
								 class="thumb-image"/>
						</a>
					</div>
					<div class="auto-info-wrap">
						<div class="auto-naziv">
							<a href="/model/audi/1213/audi-a6-2014" title="2014 Audi A6 (C7 restyle)">
								2014 Audi A6 (C7 restyle)							</a>
						</div>
						<div class="stars">
							<div class="stars5"></div>
						</div>
						<div class="euro-ncap">
							EuroNCAP:&nbsp;
							91% occupant safety						</div>
						<div>
							<b>
								<a href="/model/audi/1213/audi-a6-2014" title="2014 Audi A6 (C7 restyle)" class="orange">
							   	   	more details...						   	   </a>
					   	   </b>
						</div>
					</div>
				</div>
							<div class="auto-wrap">
					<div class="auto-slika">
						<div class="auto-score-wrap clearfix">
														<div class="score-left">
								<div class="score5b">
									<p>drivers'</p>
									<span>4.5</span>
									<p>rating</p>
								</div>
							</div>
							<div class="score-right clearfix">
								<div class="score5a"></div>
								<div class="score4a"></div>
								<div class="score3a"></div>
								<div class="score2a"></div>
								<div class="score1a"></div>
							</div>
						</div>
						<a href="/model/audi/1424/audi-a6-2018" title="">
							<img src="/resources/images/variant/2161/thumb_2018_audi_a6.jpg"
								 alt=""
								 class="thumb-image"/>
						</a>
					</div>
					<div class="auto-info-wrap">
						<div class="auto-naziv">
							<a href="/model/audi/1424/audi-a6-2018" title="2018 Audi A6 (C8)">
								2018 Audi A6 (C8)							</a>
						</div>
						<div class="stars">
							<div class="stars5"></div>
						</div>
						<div class="euro-ncap">
							EuroNCAP:&nbsp;
							93% occupant safety						</div>
						<div>
							<b>
								<a href="/model/audi/1424/audi-a6-2018" title="2018 Audi A6 (C8)" class="orange">
							   	   	more details...						   	   </a>
					   	   </b>
						</div>
					</div>
				</div>
							<div class="auto-wrap">
					<div class="auto-slika">
						<div class="auto-score-wrap clearfix">
														<div class="score-left">
								<div class="score5b">
									<p>drivers'</p>
									<span>5.0</span>
									<p>rating</p>
								</div>
							</div>
							<div class="score-right clearfix">
								<div class="score5a"></div>
								<div class="score4a"></div>
								<div class="score3a"></div>
								<div class="score2a"></div>
								<div class="score1a"></div>
							</div>
						</div>
						<a href="/model/audi/233/audi-a7-2010" title="">
							<img src="/resources/images/variant/71/thumb_a7_1.jpg"
								 alt=""
								 class="thumb-image"/>
						</a>
					</div>
					<div class="auto-info-wrap">
						<div class="auto-naziv">
							<a href="/model/audi/233/audi-a7-2010" title="2010 Audi A7 (4G)">
								2010 Audi A7 (4G)							</a>
						</div>
						<div class="stars">
							<div class="stars"></div>
						</div>
						<div class="euro-ncap">
							EuroNCAP:&nbsp;
							not tested						</div>
						<div>
							<b>
								<a href="/model/audi/233/audi-a7-2010" title="2010 Audi A7 (4G)" class="orange">
							   	   	more details...						   	   </a>
					   	   </b>
						</div>
					</div>
				</div>
							<div class="auto-wrap">
					<div class="auto-slika">
						<div class="auto-score-wrap clearfix">
														<div class="score-left">
								<div class="score5b">
									<p>drivers'</p>
									<span>4.6</span>
									<p>rating</p>
								</div>
							</div>
							<div class="score-right clearfix">
								<div class="score5a"></div>
								<div class="score4a"></div>
								<div class="score3a"></div>
								<div class="score2a"></div>
								<div class="score1a"></div>
							</div>
						</div>
						<a href="/model/audi/234/audi-a7-2014" title="">
							<img src="/resources/images/variant/72/thumb_a7_2.jpg"
								 alt=""
								 class="thumb-image"/>
						</a>
					</div>
					<div class="auto-info-wrap">
						<div class="auto-naziv">
							<a href="/model/audi/234/audi-a7-2014" title="2014 Audi A7 (4G restyle)">
								2014 Audi A7 (4G restyle)							</a>
						</div>
						<div class="stars">
							<div class="stars"></div>
						</div>
						<div class="euro-ncap">
							EuroNCAP:&nbsp;
							not tested						</div>
						<div>
							<b>
								<a href="/model/audi/234/audi-a7-2014" title="2014 Audi A7 (4G restyle)" class="orange">
							   	   	more details...						   	   </a>
					   	   </b>
						</div>
					</div>
				</div>
							<div class="auto-wrap">
					<div class="auto-slika">
						<div class="auto-score-wrap clearfix">
														<div class="score-left">
								<div class="score3b">
									<p>drivers'</p>
									<span>3.0</span>
									<p>rating</p>
								</div>
							</div>
							<div class="score-right clearfix">
								<div class="score5a"></div>
								<div class="score4a"></div>
								<div class="score3a"></div>
								<div class="score2a"></div>
								<div class="score1a"></div>
							</div>
						</div>
						<a href="/model/audi/1425/audi-a7-2018" title="">
							<img src="/resources/images/variant/2167/thumb_2018_audi_a7.jpg"
								 alt=""
								 class="thumb-image"/>
						</a>
					</div>
					<div class="auto-info-wrap">
						<div class="auto-naziv">
							<a href="/model/audi/1425/audi-a7-2018" title="2018 Audi A7 (4K)">
								2018 Audi A7 (4K)							</a>
						</div>
						<div class="stars">
							<div class="stars5"></div>
						</div>
						<div class="euro-ncap">
							EuroNCAP:&nbsp;
							93% occupant safety						</div>
						<div>
							<b>
								<a href="/model/audi/1425/audi-a7-2018" title="2018 Audi A7 (4K)" class="orange">
							   	   	more details...						   	   </a>
					   	   </b>
						</div>
					</div>
				</div>
					</div>
</div>
<div class="proizvodnja-wrap">


		<div class="proizvodnja-box variant-header accordion-controller segment-header"
			 related-accordion="accordion-F"
			 segment="F"
		 >
			<div class="proizvodnja-godina">
				F - Executive car			</div>
			<div class="proizvodnja-model">
				6 models			</div>
		</div>
		<div id="accordion-F" class="variant-accordion ">
							<div class="auto-wrap">
					<div class="auto-slika">
						<div class="auto-score-wrap clearfix">
														<div class="score-left">
								<div class="score4b">
									<p>drivers'</p>
									<span>4.2</span>
									<p>rating</p>
								</div>
							</div>
							<div class="score-right clearfix">
								<div class="score5a"></div>
								<div class="score4a"></div>
								<div class="score3a"></div>
								<div class="score2a"></div>
								<div class="score1a"></div>
							</div>
						</div>
						<a href="/model/audi/235/audi-a8-1999" title="">
							<img src="/resources/images/variant/73/thumb_a8_1.jpg"
								 alt=""
								 class="thumb-image"/>
						</a>
					</div>
					<div class="auto-info-wrap">
						<div class="auto-naziv">
							<a href="/model/audi/235/audi-a8-1999" title="1999 Audi A8 (D2 restyle)">
								1999 Audi A8 (D2 restyle)							</a>
						</div>
						<div class="stars">
							<div class="stars"></div>
						</div>
						<div class="euro-ncap">
							EuroNCAP:&nbsp;
							not tested						</div>
						<div>
							<b>
								<a href="/model/audi/235/audi-a8-1999" title="1999 Audi A8 (D2 restyle)" class="orange">
							   	   	more details...						   	   </a>
					   	   </b>
						</div>
					</div>
				</div>
							<div class="auto-wrap">
					<div class="auto-slika">
						<div class="auto-score-wrap clearfix">
														<div class="score-left">
								<div class="score5b">
									<p>drivers'</p>
									<span>4.7</span>
									<p>rating</p>
								</div>
							</div>
							<div class="score-right clearfix">
								<div class="score5a"></div>
								<div class="score4a"></div>
								<div class="score3a"></div>
								<div class="score2a"></div>
								<div class="score1a"></div>
							</div>
						</div>
						<a href="/model/audi/236/audi-a8-2002" title="">
							<img src="/resources/images/variant/74/thumb_a8_2.jpg"
								 alt=""
								 class="thumb-image"/>
						</a>
					</div>
					<div class="auto-info-wrap">
						<div class="auto-naziv">
							<a href="/model/audi/236/audi-a8-2002" title="2002 Audi A8 (D3)">
								2002 Audi A8 (D3)							</a>
						</div>
						<div class="stars">
							<div class="stars"></div>
						</div>
						<div class="euro-ncap">
							EuroNCAP:&nbsp;
							not tested						</div>
						<div>
							<b>
								<a href="/model/audi/236/audi-a8-2002" title="2002 Audi A8 (D3)" class="orange">
							   	   	more details...						   	   </a>
					   	   </b>
						</div>
					</div>
				</div>
							<div class="auto-wrap">
					<div class="auto-slika">
						<div class="auto-score-wrap clearfix">
														<div class="score-left">
								<div class="score4b">
									<p>drivers'</p>
									<span>4.3</span>
									<p>rating</p>
								</div>
							</div>
							<div class="score-right clearfix">
								<div class="score5a"></div>
								<div class="score4a"></div>
								<div class="score3a"></div>
								<div class="score2a"></div>
								<div class="score1a"></div>
							</div>
						</div>
						<a href="/model/audi/237/audi-a8-2007" title="">
							<img src="/resources/images/variant/75/thumb_a8_3.jpg"
								 alt=""
								 class="thumb-image"/>
						</a>
					</div>
					<div class="auto-info-wrap">
						<div class="auto-naziv">
							<a href="/model/audi/237/audi-a8-2007" title="2007 Audi A8 (D3 restyle)">
								2007 Audi A8 (D3 restyle)							</a>
						</div>
						<div class="stars">
							<div class="stars"></div>
						</div>
						<div class="euro-ncap">
							EuroNCAP:&nbsp;
							not tested						</div>
						<div>
							<b>
								<a href="/model/audi/237/audi-a8-2007" title="2007 Audi A8 (D3 restyle)" class="orange">
							   	   	more details...						   	   </a>
					   	   </b>
						</div>
					</div>
				</div>
							<div class="auto-wrap">
					<div class="auto-slika">
						<div class="auto-score-wrap clearfix">
														<div class="score-left">
								<div class="score5b">
									<p>drivers'</p>
									<span>4.7</span>
									<p>rating</p>
								</div>
							</div>
							<div class="score-right clearfix">
								<div class="score5a"></div>
								<div class="score4a"></div>
								<div class="score3a"></div>
								<div class="score2a"></div>
								<div class="score1a"></div>
							</div>
						</div>
						<a href="/model/audi/238/audi-a8-2010" title="">
							<img src="/resources/images/variant/76/thumb_a8_4.jpg"
								 alt=""
								 class="thumb-image"/>
						</a>
					</div>
					<div class="auto-info-wrap">
						<div class="auto-naziv">
							<a href="/model/audi/238/audi-a8-2010" title="2010 Audi A8 (D4)">
								2010 Audi A8 (D4)							</a>
						</div>
						<div class="stars">
							<div class="stars"></div>
						</div>
						<div class="euro-ncap">
							EuroNCAP:&nbsp;
							not tested						</div>
						<div>
							<b>
								<a href="/model/audi/238/audi-a8-2010" title="2010 Audi A8 (D4)" class="orange">
							   	   	more details...						   	   </a>
					   	   </b>
						</div>
					</div>
				</div>
							<div class="auto-wrap">
					<div class="auto-slika">
						<div class="auto-score-wrap clearfix">
														<div class="score-left">
								<div class="score5b">
									<p>drivers'</p>
									<span>5.0</span>
									<p>rating</p>
								</div>
							</div>
							<div class="score-right clearfix">
								<div class="score5a"></div>
								<div class="score4a"></div>
								<div class="score3a"></div>
								<div class="score2a"></div>
								<div class="score1a"></div>
							</div>
						</div>
						<a href="/model/audi/239/audi-a8-2013" title="">
							<img src="/resources/images/variant/77/thumb_a8_5.jpg"
								 alt=""
								 class="thumb-image"/>
						</a>
					</div>
					<div class="auto-info-wrap">
						<div class="auto-naziv">
							<a href="/model/audi/239/audi-a8-2013" title="2013 Audi A8 (D4 restyle)">
								2013 Audi A8 (D4 restyle)							</a>
						</div>
						<div class="stars">
							<div class="stars"></div>
						</div>
						<div class="euro-ncap">
							EuroNCAP:&nbsp;
							not tested						</div>
						<div>
							<b>
								<a href="/model/audi/239/audi-a8-2013" title="2013 Audi A8 (D4 restyle)" class="orange">
							   	   	more details...						   	   </a>
					   	   </b>
						</div>
					</div>
				</div>
							<div class="auto-wrap">
					<div class="auto-slika">
						<div class="auto-score-wrap clearfix">
														<div class="score-left">
								<div class="score3b">
									<p>drivers'</p>
									<span>3.0</span>
									<p>rating</p>
								</div>
							</div>
							<div class="score-right clearfix">
								<div class="score5a"></div>
								<div class="score4a"></div>
								<div class="score3a"></div>
								<div class="score2a"></div>
								<div class="score1a"></div>
							</div>
						</div>
						<a href="/model/audi/1454/audi-a8-2017" title="">
							<img src="/resources/images/variant/2205/thumb_2017_audi_a8.jpg"
								 alt=""
								 class="thumb-image"/>
						</a>
					</div>
					<div class="auto-info-wrap">
						<div class="auto-naziv">
							<a href="/model/audi/1454/audi-a8-2017" title="2017 Audi A8 (D5)">
								2017 Audi A8 (D5)							</a>
						</div>
						<div class="stars">
							<div class="stars"></div>
						</div>
						<div class="euro-ncap">
							EuroNCAP:&nbsp;
							not tested						</div>
						<div>
							<b>
								<a href="/model/audi/1454/audi-a8-2017" title="2017 Audi A8 (D5)" class="orange">
							   	   	more details...						   	   </a>
					   	   </b>
						</div>
					</div>
				</div>
					</div>
</div>
<div class="proizvodnja-wrap">


		<div class="proizvodnja-box variant-header accordion-controller segment-header"
			 related-accordion="accordion-S"
			 segment="S"
		 >
			<div class="proizvodnja-godina">
				S - Sports car			</div>
			<div class="proizvodnja-model">
				9 models			</div>
		</div>
		<div id="accordion-S" class="variant-accordion first-accordion">
							<div class="auto-wrap">
					<div class="auto-slika">
						<div class="auto-score-wrap clearfix">
														<div class="score-left">
								<div class="score4b">
									<p>drivers'</p>
									<span>4.5</span>
									<p>rating</p>
								</div>
							</div>
							<div class="score-right clearfix">
								<div class="score5a"></div>
								<div class="score4a"></div>
								<div class="score3a"></div>
								<div class="score2a"></div>
								<div class="score1a"></div>
							</div>
						</div>
						<a href="/model/audi/1053/audi-a5-coupe-2007" title="">
							<img src="/resources/images/variant/49/thumb_a5_1.jpg"
								 alt=""
								 class="thumb-image"/>
						</a>
					</div>
					<div class="auto-info-wrap">
						<div class="auto-naziv">
							<a href="/model/audi/1053/audi-a5-coupe-2007" title="2007 Audi A5 Coupe (8T)">
								2007 Audi A5 Coupe (8T)							</a>
						</div>
						<div class="stars">
							<div class="stars"></div>
						</div>
						<div class="euro-ncap">
							EuroNCAP:&nbsp;
							not tested						</div>
						<div>
							<b>
								<a href="/model/audi/1053/audi-a5-coupe-2007" title="2007 Audi A5 Coupe (8T)" class="orange">
							   	   	more details...						   	   </a>
					   	   </b>
						</div>
					</div>
				</div>
							<div class="auto-wrap">
					<div class="auto-slika">
						<div class="auto-score-wrap clearfix">
														<div class="score-left">
								<div class="score5b">
									<p>drivers'</p>
									<span>4.8</span>
									<p>rating</p>
								</div>
							</div>
							<div class="score-right clearfix">
								<div class="score5a"></div>
								<div class="score4a"></div>
								<div class="score3a"></div>
								<div class="score2a"></div>
								<div class="score1a"></div>
							</div>
						</div>
						<a href="/model/audi/1054/audi-a5-coupe-2011" title="">
							<img src="/resources/images/variant/51/thumb_a5_3.jpg"
								 alt=""
								 class="thumb-image"/>
						</a>
					</div>
					<div class="auto-info-wrap">
						<div class="auto-naziv">
							<a href="/model/audi/1054/audi-a5-coupe-2011" title="2011 Audi A5 Coupe (8T restyle)">
								2011 Audi A5 Coupe (8T restyle)							</a>
						</div>
						<div class="stars">
							<div class="stars"></div>
						</div>
						<div class="euro-ncap">
							EuroNCAP:&nbsp;
							not tested						</div>
						<div>
							<b>
								<a href="/model/audi/1054/audi-a5-coupe-2011" title="2011 Audi A5 Coupe (8T restyle)" class="orange">
							   	   	more details...						   	   </a>
					   	   </b>
						</div>
					</div>
				</div>
							<div class="auto-wrap">
					<div class="auto-slika">
						<div class="auto-score-wrap clearfix">
														<div class="score-left">
								<div class="score3b">
									<p>drivers'</p>
									<span>3.2</span>
									<p>rating</p>
								</div>
							</div>
							<div class="score-right clearfix">
								<div class="score5a"></div>
								<div class="score4a"></div>
								<div class="score3a"></div>
								<div class="score2a"></div>
								<div class="score1a"></div>
							</div>
						</div>
						<a href="/model/audi/1246/audi-a5-coupe-2016" title="">
							<img src="/resources/images/variant/1879/thumb_a5_8.jpg"
								 alt=""
								 class="thumb-image"/>
						</a>
					</div>
					<div class="auto-info-wrap">
						<div class="auto-naziv">
							<a href="/model/audi/1246/audi-a5-coupe-2016" title="2016 Audi A5 Coupe (F5)">
								2016 Audi A5 Coupe (F5)							</a>
						</div>
						<div class="stars">
							<div class="stars5"></div>
						</div>
						<div class="euro-ncap">
							EuroNCAP:&nbsp;
							89% occupant safety						</div>
						<div>
							<b>
								<a href="/model/audi/1246/audi-a5-coupe-2016" title="2016 Audi A5 Coupe (F5)" class="orange">
							   	   	more details...						   	   </a>
					   	   </b>
						</div>
					</div>
				</div>
							<div class="auto-wrap">
					<div class="auto-slika">
						<div class="auto-score-wrap clearfix">
														<div class="score-left">
								<div class="score3b">
									<p>drivers'</p>
									<span>3.0</span>
									<p>rating</p>
								</div>
							</div>
							<div class="score-right clearfix">
								<div class="score5a"></div>
								<div class="score4a"></div>
								<div class="score3a"></div>
								<div class="score2a"></div>
								<div class="score1a"></div>
							</div>
						</div>
						<a href="/model/audi/1340/audi-r8-2007" title="">
							<img src="/resources/images/variant/2039/thumb_r8_1.jpg"
								 alt=""
								 class="thumb-image"/>
						</a>
					</div>
					<div class="auto-info-wrap">
						<div class="auto-naziv">
							<a href="/model/audi/1340/audi-r8-2007" title="2007 Audi R8 (Typ 42)">
								2007 Audi R8 (Typ 42)							</a>
						</div>
						<div class="stars">
							<div class="stars"></div>
						</div>
						<div class="euro-ncap">
							EuroNCAP:&nbsp;
							not tested						</div>
						<div>
							<b>
								<a href="/model/audi/1340/audi-r8-2007" title="2007 Audi R8 (Typ 42)" class="orange">
							   	   	more details...						   	   </a>
					   	   </b>
						</div>
					</div>
				</div>
							<div class="auto-wrap">
					<div class="auto-slika">
						<div class="auto-score-wrap clearfix">
														<div class="score-left">
								<div class="score3b">
									<p>drivers'</p>
									<span>3.0</span>
									<p>rating</p>
								</div>
							</div>
							<div class="score-right clearfix">
								<div class="score5a"></div>
								<div class="score4a"></div>
								<div class="score3a"></div>
								<div class="score2a"></div>
								<div class="score1a"></div>
							</div>
						</div>
						<a href="/model/audi/1341/audi-r8-2012" title="">
							<img src="/resources/images/variant/2041/thumb_r8_3.jpg"
								 alt=""
								 class="thumb-image"/>
						</a>
					</div>
					<div class="auto-info-wrap">
						<div class="auto-naziv">
							<a href="/model/audi/1341/audi-r8-2012" title="2012 Audi R8 (Typ 42 restyle)">
								2012 Audi R8 (Typ 42 restyle)							</a>
						</div>
						<div class="stars">
							<div class="stars"></div>
						</div>
						<div class="euro-ncap">
							EuroNCAP:&nbsp;
							not tested						</div>
						<div>
							<b>
								<a href="/model/audi/1341/audi-r8-2012" title="2012 Audi R8 (Typ 42 restyle)" class="orange">
							   	   	more details...						   	   </a>
					   	   </b>
						</div>
					</div>
				</div>
							<div class="auto-wrap">
					<div class="auto-slika">
						<div class="auto-score-wrap clearfix">
														<div class="score-left">
								<div class="score4b">
									<p>drivers'</p>
									<span>4.0</span>
									<p>rating</p>
								</div>
							</div>
							<div class="score-right clearfix">
								<div class="score5a"></div>
								<div class="score4a"></div>
								<div class="score3a"></div>
								<div class="score2a"></div>
								<div class="score1a"></div>
							</div>
						</div>
						<a href="/model/audi/240/audi-tt-1998" title="">
							<img src="/resources/images/variant/84/thumb_tt_1.jpg"
								 alt=""
								 class="thumb-image"/>
						</a>
					</div>
					<div class="auto-info-wrap">
						<div class="auto-naziv">
							<a href="/model/audi/240/audi-tt-1998" title="1998 Audi TT (8N)">
								1998 Audi TT (8N)							</a>
						</div>
						<div class="stars">
							<div class="stars4"></div>
						</div>
						<div class="euro-ncap">
							EuroNCAP:&nbsp;
							~70% occupant safety						</div>
						<div>
							<b>
								<a href="/model/audi/240/audi-tt-1998" title="1998 Audi TT (8N)" class="orange">
							   	   	more details...						   	   </a>
					   	   </b>
						</div>
					</div>
				</div>
							<div class="auto-wrap">
					<div class="auto-slika">
						<div class="auto-score-wrap clearfix">
														<div class="score-left">
								<div class="score5b">
									<p>drivers'</p>
									<span>4.6</span>
									<p>rating</p>
								</div>
							</div>
							<div class="score-right clearfix">
								<div class="score5a"></div>
								<div class="score4a"></div>
								<div class="score3a"></div>
								<div class="score2a"></div>
								<div class="score1a"></div>
							</div>
						</div>
						<a href="/model/audi/241/audi-tt-2006" title="">
							<img src="/resources/images/variant/86/thumb_tt_3.jpg"
								 alt=""
								 class="thumb-image"/>
						</a>
					</div>
					<div class="auto-info-wrap">
						<div class="auto-naziv">
							<a href="/model/audi/241/audi-tt-2006" title="2006 Audi TT (8J)">
								2006 Audi TT (8J)							</a>
						</div>
						<div class="stars">
							<div class="stars4"></div>
						</div>
						<div class="euro-ncap">
							EuroNCAP:&nbsp;
							~70% occupant safety						</div>
						<div>
							<b>
								<a href="/model/audi/241/audi-tt-2006" title="2006 Audi TT (8J)" class="orange">
							   	   	more details...						   	   </a>
					   	   </b>
						</div>
					</div>
				</div>
							<div class="auto-wrap">
					<div class="auto-slika">
						<div class="auto-score-wrap clearfix">
														<div class="score-left">
								<div class="score5b">
									<p>drivers'</p>
									<span>5.0</span>
									<p>rating</p>
								</div>
							</div>
							<div class="score-right clearfix">
								<div class="score5a"></div>
								<div class="score4a"></div>
								<div class="score3a"></div>
								<div class="score2a"></div>
								<div class="score1a"></div>
							</div>
						</div>
						<a href="/model/audi/242/audi-tt-2010" title="">
							<img src="/resources/images/variant/88/thumb_tt_5.jpg"
								 alt=""
								 class="thumb-image"/>
						</a>
					</div>
					<div class="auto-info-wrap">
						<div class="auto-naziv">
							<a href="/model/audi/242/audi-tt-2010" title="2010 Audi TT (8J restyle)">
								2010 Audi TT (8J restyle)							</a>
						</div>
						<div class="stars">
							<div class="stars4"></div>
						</div>
						<div class="euro-ncap">
							EuroNCAP:&nbsp;
							~70% occupant safety						</div>
						<div>
							<b>
								<a href="/model/audi/242/audi-tt-2010" title="2010 Audi TT (8J restyle)" class="orange">
							   	   	more details...						   	   </a>
					   	   </b>
						</div>
					</div>
				</div>
							<div class="auto-wrap">
					<div class="auto-slika">
						<div class="auto-score-wrap clearfix">
														<div class="score-left">
								<div class="score5b">
									<p>drivers'</p>
									<span>5.0</span>
									<p>rating</p>
								</div>
							</div>
							<div class="score-right clearfix">
								<div class="score5a"></div>
								<div class="score4a"></div>
								<div class="score3a"></div>
								<div class="score2a"></div>
								<div class="score1a"></div>
							</div>
						</div>
						<a href="/model/audi/1057/audi-tt-2015" title="">
							<img src="/resources/images/variant/280/thumb_tt_7.jpg"
								 alt=""
								 class="thumb-image"/>
						</a>
					</div>
					<div class="auto-info-wrap">
						<div class="auto-naziv">
							<a href="/model/audi/1057/audi-tt-2015" title="2015 Audi TT (FV)">
								2015 Audi TT (FV)							</a>
						</div>
						<div class="stars">
							<div class="stars4"></div>
						</div>
						<div class="euro-ncap">
							EuroNCAP:&nbsp;
							81% occupant safety						</div>
						<div>
							<b>
								<a href="/model/audi/1057/audi-tt-2015" title="2015 Audi TT (FV)" class="orange">
							   	   	more details...						   	   </a>
					   	   </b>
						</div>
					</div>
				</div>
					</div>
</div>
<div class="proizvodnja-wrap">


		<div class="proizvodnja-box variant-header accordion-controller segment-header"
			 related-accordion="accordion-M"
			 segment="M"
		 >
			<div class="proizvodnja-godina">
				M - MPV			</div>
			<div class="proizvodnja-model">
				1 model			</div>
		</div>
		<div id="accordion-M" class="variant-accordion ">
							<div class="auto-wrap">
					<div class="auto-slika">
						<div class="auto-score-wrap clearfix">
														<div class="score-left">
								<div class="score4b">
									<p>drivers'</p>
									<span>4.3</span>
									<p>rating</p>
								</div>
							</div>
							<div class="score-right clearfix">
								<div class="score5a"></div>
								<div class="score4a"></div>
								<div class="score3a"></div>
								<div class="score2a"></div>
								<div class="score1a"></div>
							</div>
						</div>
						<a href="/model/audi/213/audi-a2-2000" title="">
							<img src="/resources/images/variant/24/thumb_a2.jpg"
								 alt=""
								 class="thumb-image"/>
						</a>
					</div>
					<div class="auto-info-wrap">
						<div class="auto-naziv">
							<a href="/model/audi/213/audi-a2-2000" title="2000 Audi A2 (8Z)">
								2000 Audi A2 (8Z)							</a>
						</div>
						<div class="stars">
							<div class="stars4"></div>
						</div>
						<div class="euro-ncap">
							EuroNCAP:&nbsp;
							~70% occupant safety						</div>
						<div>
							<b>
								<a href="/model/audi/213/audi-a2-2000" title="2000 Audi A2 (8Z)" class="orange">
							   	   	more details...						   	   </a>
					   	   </b>
						</div>
					</div>
				</div>
					</div>
</div>
<div class="proizvodnja-wrap">


		<div class="proizvodnja-box variant-header accordion-controller segment-header"
			 related-accordion="accordion-J"
			 segment="J"
		 >
			<div class="proizvodnja-godina">
				J - SUV			</div>
			<div class="proizvodnja-model">
				11 models			</div>
		</div>
		<div id="accordion-J" class="variant-accordion ">
							<div class="auto-wrap">
					<div class="auto-slika">
						<div class="auto-score-wrap clearfix">
														<div class="score-left">
								<div class="score3b">
									<p>drivers'</p>
									<span>3.0</span>
									<p>rating</p>
								</div>
							</div>
							<div class="score-right clearfix">
								<div class="score5a"></div>
								<div class="score4a"></div>
								<div class="score3a"></div>
								<div class="score2a"></div>
								<div class="score1a"></div>
							</div>
						</div>
						<a href="/model/audi/1291/audi-q2-2016" title="">
							<img src="/resources/images/variant/1949/thumb_q2_1.jpg"
								 alt=""
								 class="thumb-image"/>
						</a>
					</div>
					<div class="auto-info-wrap">
						<div class="auto-naziv">
							<a href="/model/audi/1291/audi-q2-2016" title="2016 Audi Q2 (GA)">
								2016 Audi Q2 (GA)							</a>
						</div>
						<div class="stars">
							<div class="stars5"></div>
						</div>
						<div class="euro-ncap">
							EuroNCAP:&nbsp;
							93% occupant safety						</div>
						<div>
							<b>
								<a href="/model/audi/1291/audi-q2-2016" title="2016 Audi Q2 (GA)" class="orange">
							   	   	more details...						   	   </a>
					   	   </b>
						</div>
					</div>
				</div>
							<div class="auto-wrap">
					<div class="auto-slika">
						<div class="auto-score-wrap clearfix">
														<div class="score-left">
								<div class="score5b">
									<p>drivers'</p>
									<span>5.0</span>
									<p>rating</p>
								</div>
							</div>
							<div class="score-right clearfix">
								<div class="score5a"></div>
								<div class="score4a"></div>
								<div class="score3a"></div>
								<div class="score2a"></div>
								<div class="score1a"></div>
							</div>
						</div>
						<a href="/model/audi/243/audi-q3-2011" title="">
							<img src="/resources/images/variant/78/thumb_q3_1.jpg"
								 alt=""
								 class="thumb-image"/>
						</a>
					</div>
					<div class="auto-info-wrap">
						<div class="auto-naziv">
							<a href="/model/audi/243/audi-q3-2011" title="2011 Audi Q3 (8U)">
								2011 Audi Q3 (8U)							</a>
						</div>
						<div class="stars">
							<div class="stars5"></div>
						</div>
						<div class="euro-ncap">
							EuroNCAP:&nbsp;
							94% occupant safety						</div>
						<div>
							<b>
								<a href="/model/audi/243/audi-q3-2011" title="2011 Audi Q3 (8U)" class="orange">
							   	   	more details...						   	   </a>
					   	   </b>
						</div>
					</div>
				</div>
							<div class="auto-wrap">
					<div class="auto-slika">
						<div class="auto-score-wrap clearfix">
														<div class="score-left">
								<div class="score5b">
									<p>drivers'</p>
									<span>5.0</span>
									<p>rating</p>
								</div>
							</div>
							<div class="score-right clearfix">
								<div class="score5a"></div>
								<div class="score4a"></div>
								<div class="score3a"></div>
								<div class="score2a"></div>
								<div class="score1a"></div>
							</div>
						</div>
						<a href="/model/audi/1058/audi-q3-2015" title="">
							<img src="/resources/images/variant/281/thumb_q3_2.jpg"
								 alt=""
								 class="thumb-image"/>
						</a>
					</div>
					<div class="auto-info-wrap">
						<div class="auto-naziv">
							<a href="/model/audi/1058/audi-q3-2015" title="2015 Audi Q3 (8U restyle)">
								2015 Audi Q3 (8U restyle)							</a>
						</div>
						<div class="stars">
							<div class="stars5"></div>
						</div>
						<div class="euro-ncap">
							EuroNCAP:&nbsp;
							94% occupant safety						</div>
						<div>
							<b>
								<a href="/model/audi/1058/audi-q3-2015" title="2015 Audi Q3 (8U restyle)" class="orange">
							   	   	more details...						   	   </a>
					   	   </b>
						</div>
					</div>
				</div>
							<div class="auto-wrap">
					<div class="auto-slika">
						<div class="auto-score-wrap clearfix">
														<div class="score-left">
								<div class="score3b">
									<p>drivers'</p>
									<span>3.0</span>
									<p>rating</p>
								</div>
							</div>
							<div class="score-right clearfix">
								<div class="score5a"></div>
								<div class="score4a"></div>
								<div class="score3a"></div>
								<div class="score2a"></div>
								<div class="score1a"></div>
							</div>
						</div>
						<a href="/model/audi/1414/audi-q3-2018" title="">
							<img src="/resources/images/variant/2144/thumb_2019_audi_q3.jpg"
								 alt=""
								 class="thumb-image"/>
						</a>
					</div>
					<div class="auto-info-wrap">
						<div class="auto-naziv">
							<a href="/model/audi/1414/audi-q3-2018" title="2018 Audi Q3 (F3)">
								2018 Audi Q3 (F3)							</a>
						</div>
						<div class="stars">
							<div class="stars5"></div>
						</div>
						<div class="euro-ncap">
							EuroNCAP:&nbsp;
							95% occupant safety						</div>
						<div>
							<b>
								<a href="/model/audi/1414/audi-q3-2018" title="2018 Audi Q3 (F3)" class="orange">
							   	   	more details...						   	   </a>
					   	   </b>
						</div>
					</div>
				</div>
							<div class="auto-wrap">
					<div class="auto-slika">
						<div class="auto-score-wrap clearfix">
														<div class="score-left">
								<div class="score5b">
									<p>drivers'</p>
									<span>5.0</span>
									<p>rating</p>
								</div>
							</div>
							<div class="score-right clearfix">
								<div class="score5a"></div>
								<div class="score4a"></div>
								<div class="score3a"></div>
								<div class="score2a"></div>
								<div class="score1a"></div>
							</div>
						</div>
						<a href="/model/audi/244/audi-q5-2008" title="">
							<img src="/resources/images/variant/79/thumb_q5_1.jpg"
								 alt=""
								 class="thumb-image"/>
						</a>
					</div>
					<div class="auto-info-wrap">
						<div class="auto-naziv">
							<a href="/model/audi/244/audi-q5-2008" title="2008 Audi Q5 (8R)">
								2008 Audi Q5 (8R)							</a>
						</div>
						<div class="stars">
							<div class="stars5"></div>
						</div>
						<div class="euro-ncap">
							EuroNCAP:&nbsp;
							92% occupant safety						</div>
						<div>
							<b>
								<a href="/model/audi/244/audi-q5-2008" title="2008 Audi Q5 (8R)" class="orange">
							   	   	more details...						   	   </a>
					   	   </b>
						</div>
					</div>
				</div>
							<div class="auto-wrap">
					<div class="auto-slika">
						<div class="auto-score-wrap clearfix">
														<div class="score-left">
								<div class="score4b">
									<p>drivers'</p>
									<span>4.3</span>
									<p>rating</p>
								</div>
							</div>
							<div class="score-right clearfix">
								<div class="score5a"></div>
								<div class="score4a"></div>
								<div class="score3a"></div>
								<div class="score2a"></div>
								<div class="score1a"></div>
							</div>
						</div>
						<a href="/model/audi/245/audi-q5-2012" title="">
							<img src="/resources/images/variant/80/thumb_q5_2.jpg"
								 alt=""
								 class="thumb-image"/>
						</a>
					</div>
					<div class="auto-info-wrap">
						<div class="auto-naziv">
							<a href="/model/audi/245/audi-q5-2012" title="2012 Audi Q5 (8R restyle)">
								2012 Audi Q5 (8R restyle)							</a>
						</div>
						<div class="stars">
							<div class="stars5"></div>
						</div>
						<div class="euro-ncap">
							EuroNCAP:&nbsp;
							92% occupant safety						</div>
						<div>
							<b>
								<a href="/model/audi/245/audi-q5-2012" title="2012 Audi Q5 (8R restyle)" class="orange">
							   	   	more details...						   	   </a>
					   	   </b>
						</div>
					</div>
				</div>
							<div class="auto-wrap">
					<div class="auto-slika">
						<div class="auto-score-wrap clearfix">
														<div class="score-left">
								<div class="score5b">
									<p>drivers'</p>
									<span>4.5</span>
									<p>rating</p>
								</div>
							</div>
							<div class="score-right clearfix">
								<div class="score5a"></div>
								<div class="score4a"></div>
								<div class="score3a"></div>
								<div class="score2a"></div>
								<div class="score1a"></div>
							</div>
						</div>
						<a href="/model/audi/1304/audi-q5-2016" title="">
							<img src="/resources/images/variant/1972/thumb_q5_3.jpg"
								 alt=""
								 class="thumb-image"/>
						</a>
					</div>
					<div class="auto-info-wrap">
						<div class="auto-naziv">
							<a href="/model/audi/1304/audi-q5-2016" title="2016 Audi Q5 (FY)">
								2016 Audi Q5 (FY)							</a>
						</div>
						<div class="stars">
							<div class="stars5"></div>
						</div>
						<div class="euro-ncap">
							EuroNCAP:&nbsp;
							93% occupant safety						</div>
						<div>
							<b>
								<a href="/model/audi/1304/audi-q5-2016" title="2016 Audi Q5 (FY)" class="orange">
							   	   	more details...						   	   </a>
					   	   </b>
						</div>
					</div>
				</div>
							<div class="auto-wrap">
					<div class="auto-slika">
						<div class="auto-score-wrap clearfix">
														<div class="score-left">
								<div class="score5b">
									<p>drivers'</p>
									<span>4.9</span>
									<p>rating</p>
								</div>
							</div>
							<div class="score-right clearfix">
								<div class="score5a"></div>
								<div class="score4a"></div>
								<div class="score3a"></div>
								<div class="score2a"></div>
								<div class="score1a"></div>
							</div>
						</div>
						<a href="/model/audi/246/audi-q7-2006" title="">
							<img src="/resources/images/variant/81/thumb_q7_1.jpg"
								 alt=""
								 class="thumb-image"/>
						</a>
					</div>
					<div class="auto-info-wrap">
						<div class="auto-naziv">
							<a href="/model/audi/246/audi-q7-2006" title="2006 Audi Q7 (4L)">
								2006 Audi Q7 (4L)							</a>
						</div>
						<div class="stars">
							<div class="stars4"></div>
						</div>
						<div class="euro-ncap">
							EuroNCAP:&nbsp;
							~70% occupant safety						</div>
						<div>
							<b>
								<a href="/model/audi/246/audi-q7-2006" title="2006 Audi Q7 (4L)" class="orange">
							   	   	more details...						   	   </a>
					   	   </b>
						</div>
					</div>
				</div>
							<div class="auto-wrap">
					<div class="auto-slika">
						<div class="auto-score-wrap clearfix">
														<div class="score-left">
								<div class="score5b">
									<p>drivers'</p>
									<span>5.0</span>
									<p>rating</p>
								</div>
							</div>
							<div class="score-right clearfix">
								<div class="score5a"></div>
								<div class="score4a"></div>
								<div class="score3a"></div>
								<div class="score2a"></div>
								<div class="score1a"></div>
							</div>
						</div>
						<a href="/model/audi/247/audi-q7-2009" title="">
							<img src="/resources/images/variant/82/thumb_q7_2.jpg"
								 alt=""
								 class="thumb-image"/>
						</a>
					</div>
					<div class="auto-info-wrap">
						<div class="auto-naziv">
							<a href="/model/audi/247/audi-q7-2009" title="2009 Audi Q7 (4L restyle)">
								2009 Audi Q7 (4L restyle)							</a>
						</div>
						<div class="stars">
							<div class="stars4"></div>
						</div>
						<div class="euro-ncap">
							EuroNCAP:&nbsp;
							~70% occupant safety						</div>
						<div>
							<b>
								<a href="/model/audi/247/audi-q7-2009" title="2009 Audi Q7 (4L restyle)" class="orange">
							   	   	more details...						   	   </a>
					   	   </b>
						</div>
					</div>
				</div>
							<div class="auto-wrap">
					<div class="auto-slika">
						<div class="auto-score-wrap clearfix">
														<div class="score-left">
								<div class="score5b">
									<p>drivers'</p>
									<span>4.7</span>
									<p>rating</p>
								</div>
							</div>
							<div class="score-right clearfix">
								<div class="score5a"></div>
								<div class="score4a"></div>
								<div class="score3a"></div>
								<div class="score2a"></div>
								<div class="score1a"></div>
							</div>
						</div>
						<a href="/model/audi/1055/audi-q7-2015" title="">
							<img src="/resources/images/variant/83/thumb_q7_3.jpg"
								 alt=""
								 class="thumb-image"/>
						</a>
					</div>
					<div class="auto-info-wrap">
						<div class="auto-naziv">
							<a href="/model/audi/1055/audi-q7-2015" title="2015 Audi Q7 (4M)">
								2015 Audi Q7 (4M)							</a>
						</div>
						<div class="stars">
							<div class="stars5"></div>
						</div>
						<div class="euro-ncap">
							EuroNCAP:&nbsp;
							94% occupant safety						</div>
						<div>
							<b>
								<a href="/model/audi/1055/audi-q7-2015" title="2015 Audi Q7 (4M)" class="orange">
							   	   	more details...						   	   </a>
					   	   </b>
						</div>
					</div>
				</div>
							<div class="auto-wrap">
					<div class="auto-slika">
						<div class="auto-score-wrap clearfix">
														<div class="score-left">
								<div class="score3b">
									<p>drivers'</p>
									<span>3.0</span>
									<p>rating</p>
								</div>
							</div>
							<div class="score-right clearfix">
								<div class="score5a"></div>
								<div class="score4a"></div>
								<div class="score3a"></div>
								<div class="score2a"></div>
								<div class="score1a"></div>
							</div>
						</div>
						<a href="/model/audi/1430/audi-q8-2018" title="">
							<img src="/resources/images/variant/2175/thumb_2018_audi_q8.jpg"
								 alt=""
								 class="thumb-image"/>
						</a>
					</div>
					<div class="auto-info-wrap">
						<div class="auto-naziv">
							<a href="/model/audi/1430/audi-q8-2018" title="2018 Audi Q8 (4M)">
								2018 Audi Q8 (4M)							</a>
						</div>
						<div class="stars">
							<div class="stars5"></div>
						</div>
						<div class="euro-ncap">
							EuroNCAP:&nbsp;
							93% occupant safety						</div>
						<div>
							<b>
								<a href="/model/audi/1430/audi-q8-2018" title="2018 Audi Q8 (4M)" class="orange">
							   	   	more details...						   	   </a>
					   	   </b>
						</div>
					</div>
				</div>
					</div>
</div>

</div>
<div>
	<h2><i class="fa fa-search" aria-hidden="true"></i>Check a car by its VIN number</h2>
</div>
<div class="box-wrap" style="height: 250px;">
	<div
	  data-cvaff
	  data-locale="en-EU"
	  data-a="automanijak"
	  data-b="81ec5429"
	  style="width: 100%; height: 100%">
	</div>
</div>
	</div>
			</div>
		</div>
		<div id="footer">
			<div class="footer-wrap">
				Copyright &copy; 2015 - 2022 automaniac.org - All rights reserved.
				</br>
				Powered by <strong>Gama1 Solutions</strong>
				<div class="footer-links">
					<div class="footer-menu">
						<ul>
							<li><a href="/">HOME page</a></li>
							<li><a href="/about">about</a></li>
							<li><a href="/terms">terms & conditions</a></li>
							<li><a href="/legal">legal note</a></li>
							<li><a href="/faq">frequent questions (FAQ)</a></li>
							<li><a href="/contact">contact us</a></li>
						</ul>
					</div>

					<div class='footer-auto-brend'><ul><li><a href='/make/alfa-romeo'>Alfa Romeo</a></li><li><a href='/make/audi'>Audi</a></li><li><a href='/make/bmw'>BMW</a></li><li><a href='/make/chery'>Chery</a></li><li><a href='/make/chevrolet'>Chevrolet</a></li><li><a href='/make/chrysler'>Chrysler</a></li></ul></div><div class='footer-auto-brend'><ul><li><a href='/make/citroen'>Citroen</a></li><li><a href='/make/dacia'>Dacia</a></li><li><a href='/make/daihatsu'>Daihatsu</a></li><li><a href='/make/dodge'>Dodge</a></li><li><a href='/make/fiat'>FIAT</a></li><li><a href='/make/ford'>Ford</a></li></ul></div><div class='footer-auto-brend'><ul><li><a href='/make/honda'>Honda</a></li><li><a href='/make/hyundai'>Hyundai</a></li><li><a href='/make/infiniti'>Infiniti</a></li><li><a href='/make/jaguar'>Jaguar</a></li><li><a href='/make/jeep'>Jeep</a></li><li><a href='/make/kia'>KIA</a></li></ul></div><div class='footer-auto-brend'><ul><li><a href='/make/lada'>Lada</a></li><li><a href='/make/lancia'>Lancia</a></li><li><a href='/make/land-rover'>Land Rover</a></li><li><a href='/make/lexus'>Lexus</a></li><li><a href='/make/mazda'>Mazda</a></li><li><a href='/make/mercedes-benz'>Mercedes</a></li></ul></div><div class='footer-auto-brend'><ul><li><a href='/make/mini'>Mini</a></li><li><a href='/make/mitsubishi'>Mitsubishi</a></li><li><a href='/make/nissan'>Nissan</a></li><li><a href='/make/opel'>Opel</a></li><li><a href='/make/peugeot'>Peugeot</a></li><li><a href='/make/porsche'>Porsche</a></li></ul></div><div class='footer-auto-brend'><ul><li><a href='/make/renault'>Renault</a></li><li><a href='/make/rover'>Rover</a></li><li><a href='/make/saab'>SAAB</a></li><li><a href='/make/ssangyong'>SSangYong</a></li><li><a href='/make/seat'>Seat</a></li><li><a href='/make/smart'>Smart</a></li></ul></div><div class='footer-auto-brend'><ul><li><a href='/make/subaru'>Subaru</a></li><li><a href='/make/suzuki'>Suzuki</a></li><li><a href='/make/tesla'>Tesla</a></li><li><a href='/make/toyota'>Toyota</a></li><li><a href='/make/volkswagen'>Volkswagen</a></li><li><a href='/make/volvo'>Volvo</a></li></ul></div><div class='footer-auto-brend'><ul><li><a href='/make/zastava'>Zastava</a></li><li><a href='/make/škoda'>Škoda</a></li></ul></div>
				</div>
			</div>
		</div>
	</body>
</html>`

const variants_html = `

<!DOCTYPE html>
<html lang="en">
	<head>
		<meta charset="utf-8">
		<title>
			Audi A4 (2018 - ) - AutoManiac		</title>
		<meta content="width=device-width, initial-scale=1.0" name="viewport">
		<meta content="Audi A4 specifications,Audi A4 consumption,Audi A4 crash test,Audi A4 prices,Audi A4 buying,Audi A4 selling,Audi A4 opinions" name="keywords">
		<meta content="2018 Audi A4 - Information, specifications and drivers opinions" name="description">
		<meta http-equiv="Cache-Control" content="no-cache, max-age=0, must-revalidate, no-store">
		<meta property="og:site_name" content="AutoManiac"/>
		<meta property="og:image" content="/resources/images/model/1394/2018_audi_a4_interior.jpg" />
		<meta property="og:image:width" content="700"/>
		<meta property="og:image:height" content="300"/>
		<meta property="og:description" content="2018 Audi A4 - Information, specifications and drivers opinions" />
		<meta property="og:title" content="Audi A4 (2018 - )" />
		<meta property="fb:app_id" content="290604621342699" />
		<meta property="fb:admins" content="1276696339"/>
		<link rel="image_src" type="image/jpeg" href="/resources/images/model/1394/2018_audi_a4_interior.jpg" />
		<link rel="shortcut icon" href="/resources/images/gui/favicon.png" type="image/x-icon">
		<link rel="icon" href="/resources/images/gui/favicon.png" type="image/x-icon">
			<script type="text/javascript" src="/resources/js/jquery-1.8.3.js"></script>
	<script type="text/javascript" src="/resources/js/jquery-ui.js"></script>
	<script type="text/javascript" src="/resources/js/scripts.js?version=v2.1.0"></script>
	<script type="text/javascript" src="/resources/js/jquery.chosen.min.js"></script>
	<script src="https://apis.google.com/js/platform.js" async defer></script>
	<script async src="https://securepubads.g.doubleclick.net/tag/js/gpt.js"></script>
	<!--
	<script async src="//pagead2.googlesyndication.com/pagead/js/adsbygoogle.js"></script>
	<script type="text/javascript" src="/resources/js/adsense.js"></script>
	-->
	<script async defer crossorigin="anonymous" src="https://connect.facebook.net/en_GB/sdk.js#xfbml=1&version=v4.0"></script>
	
						
	<link rel="stylesheet" href="/resources/css/styles.css?version=v2.1.6" type="text/css" media="screen" charset="utf-8" />
	<link rel="stylesheet" href="/resources/css/chosen.css" type="text/css" media="screen" charset="utf-8" />
	<link rel="stylesheet" href="/resources/css/font-awesome.min.css" >
	<link rel="stylesheet" href="https://fonts.googleapis.com/css?family=Open+Sans:300,300i,400,400i,700,700i|Raleway:300,400,500,700,800|Montserrat:300,400,700">
		<link rel="icon" href="/resources/images/gui/favicon.png" >
		<link rel="shortcut icon" href="/resources/images/gui/favicon.png" type="image/x-icon">
		<link rel="apple-touch-icon" href="images/gui/apple-touch-icon.png">
	</head>

	<body root-url="">

		<!-- Load Facebook SDK for JavaScript -->
		<div id="fb-root"></div>
		<script>
		  window.fbAsyncInit = function() {
		    FB.init({
		      appId      : '290604621342699',
		      xfbml      : true,
		      version    : 'v2.8'
		    });
		  };

		  (function(d, s, id){
		     var js, fjs = d.getElementsByTagName(s)[0];
		     if (d.getElementById(id)) {return;}
		     js = d.createElement(s); js.id = id;
		     js.src = "//connect.facebook.net/en_US/sdk.js";
		     fjs.parentNode.insertBefore(js, fjs);
		   }(document, 'script', 'facebook-jssdk'));

		  (function(w,d,u,h,s){
		    h=d.getElementsByTagName('head')[0];
		    s=d.createElement('script');
		    s.async=1;
		    s.src=u+'/sdk.js';
		    h.appendChild(s);
		  })(window,document,'https://aff.carvertical.com');

			
			  (function(i,s,o,g,r,a,m){i['GoogleAnalyticsObject']=r;i[r]=i[r]||function(){
			  (i[r].q=i[r].q||[]).push(arguments)},i[r].l=1*new Date();a=s.createElement(o),
			  m=s.getElementsByTagName(o)[0];a.async=1;a.src=g;m.parentNode.insertBefore(a,m)
			  })(window,document,'script','//www.google-analytics.com/analytics.js','ga');
			
			  ga('create', 'UA-75322623-2', 'auto');
	  		  ga('send', 'pageview');

	  		
		
		</script>

		<div class="menu-wrap">
			<div class="topnav" id="myTopnav">
				<div class="top-bar">
					<div class="top-bar-wrap">
						<div class="top-bar-left">
							<a href="/about" title="about">about</a>
							<a href="/legal" title="legal note">legal note</a>
							<a href="/contact" title="contact us">contact us</a>
							<span style="color: #fff; font-weight: bold;"></span>
						</div>
						<div class="top-bar-right">
							<a href="#" title="">
								Login <i class="fa fa-user" aria-hidden="true"></i>
							</a>
							<a href="https://www.automanie.org" title="">
								deutsch (D, AT, CH)<i class="fa fa-globe" aria-hidden="true"></i>
							</a>
							<a href="https://www.automanijak.com" title="">
								srpski (ex-yugoslavia)<i class="fa fa-flag" aria-hidden="true"></i>
							</a>
						</div>
					</div>
				</div>
				<a href="#" class="icon"><i class="fa fa-bars"></i></a>
				<div class="bottom-bar">
					<div class="bottom-bar-wrap">
						<div class="bottom-bar-left">
							<div class="logo">
								<a href="/" title="">
									<img src="/resources/images/gui/logo.png" alt="" title="">
								</a>
							</div>
						</div>
						<div class="bottom-bar-right">
							<a href="/virtual-adviser">Virtual adviser</a>
							<a href="/specs">Specifications by model</a>
							<a href="/compare">Compare two cars</a>
							<a href="/badges">Automotive badges</a>
							<a href="/blog">Blog</a>
						</div>
					</div>
				</div>
			</div>
		</div>
		<div id="main-wrapper">
			<div class="left-ad">
				<script async src="//pagead2.googlesyndication.com/pagead/js/adsbygoogle.js"></script>
<!-- branding -->
<ins class="adsbygoogle branding-ad"
     style="width:300px;height:600px"
     data-ad-client="ca-pub-2798966384669812"
     data-ad-slot="1676022217"></ins>
<script>
(adsbygoogle = window.adsbygoogle || []).push({});
</script>				<script async src="https://pagead2.googlesyndication.com/pagead/js/adsbygoogle.js"></script>
<!-- massive-header -->
<ins class="adsbygoogle massive-header-ad"
     data-ad-client="ca-pub-2798966384669812"
     data-ad-slot="9878901748"
     data-full-width-responsive="true"></ins>
<script>
     (adsbygoogle = window.adsbygoogle || []).push({});
</script>			</div>
			<div class="right-ad">
				<script async src="//pagead2.googlesyndication.com/pagead/js/adsbygoogle.js"></script>
<!-- branding -->
<ins class="adsbygoogle branding-ad"
     style="width:300px;height:600px"
     data-ad-client="ca-pub-2798966384669812"
     data-ad-slot="1676022217"></ins>
<script>
(adsbygoogle = window.adsbygoogle || []).push({});
</script>			</div>
			<div id="wrapper" class="clearfix">	<div id="sidebar">
		<div class="sidebar-wrap">
		<div id="sidebar-dodatno" class="search-box">
	<div class="sidebar-dodatno-wrap">
		<script async src="https://cse.google.com/cse.js?cx=016645478549947938819:d5ilzoxfr9e"></script>
		<script>
			window.onload = function(){
				$(".gsc-input").attr("placeholder", "Search this website...");
			}
		</script>
		<div class="gcse-search"></div>
	</div>
</div>	
	<div id="sidebar-dodatno">
		<div class="sidebar-misljenje-naslov">
			<h2>Drivers' opinion</h2>
		</div>
		<div class="sidebar-dodatno-wrap">
		<div class="misljenje-vozaca">
	<span>Reliability <strong title="- Malfunction frequency&#13;- Malfunction severity&#13;- Spare parts availability">(?)</strong>:</span>
	<div class="stars">
		<div class="stars0"></div>
	</div>
	<div class="misljenje-ocena">?</div>
</div><div class="misljenje-vozaca">
	<span>Performance <strong title="- Acceleration&#13;- Top speed&#13;- Braking efficiency&#13;- Handling">(?)</strong>:</span>
	<div class="stars">
		<div class="stars0"></div>
	</div>
	<div class="misljenje-ocena">?</div>
</div><div class="misljenje-vozaca">
	<span>Economy <strong title="- Fuel consumption&#13;- Regular maintenance prices&#13;- Spare parts prices&#13;- Other expenses specific to this car">(?)</strong>:</span>
	<div class="stars">
		<div class="stars0"></div>
	</div>
	<div class="misljenje-ocena">?</div>
</div><div class="misljenje-vozaca">
	<span>Comfort <strong title="- Suspension quality&#13;- Seats comfort&#13;- Legs and head space&#13;- Driver's position & ergonomy">(?)</strong>:</span>
	<div class="stars">
		<div class="stars0"></div>
	</div>
	<div class="misljenje-ocena">?</div>
</div><div class="misljenje-vozaca">
	<span>Practicality <strong title="- Easy to use&#13;- Quantity & quality of in-car compartments&#13;- Luggage space access&#13;- Spare parts replacement ease">(?)</strong>:</span>
	<div class="stars">
		<div class="stars0"></div>
	</div>
	<div class="misljenje-ocena">?</div>
</div>		</div>
	</div><div id="sidebar-dodatno">
	<div class="sidebar-minimalna-cena-naslov">
		<h2>Starting price</h2>
	</div>
	<div class="sidebar-dodatno-wrap">
		<div class="dodatno-text">
			<div>
				from				</br>
				<strong>22.900</strong>
				</br>
				EUR*
			</div>
			<span>Starting price is informative only and is being updated using information from leading online yellow pages. Prices of well maintained examples are commonly 20% higher.</span> 
		</div>
	</div>
</div>	
	<div class="najpopulariji-wrap-bottom">
		<div class="sidebar-najpopularniji-naslov">
			<h2>Some similar models...</h2>
		</div>
		<div class="najpop-item-wrap">
			
	<div class="najpop-item">
		<div class="najpop-item-name">
			<a href="/model/bmw/1219/bmw-3-series-2015" title="2015 BMW 3 Series (F30 rest...">
				2015 BMW 3 Series (F30 rest...			</a>
		</div>
		<div class="najpop-item-img">
			<a href="/model/bmw/1219/bmw-3-series-2015" title="2015 BMW 3 Series (F30 rest...">
				<img src="/resources/images/variant/1827/thumb_3_21.jpg" 	title="BMW 3 Series"
					 alt="BMW 3 Series"
				 	 class="thumb-image">
			</a>
			<div class="ocena"><span>4.6</span>
			</div>
		</div>
		<div class="najpop-item-val">
			from: 			</br>
			<strong>7.900</strong>
			</br>EUR</div>
	</div>	
	<div class="najpop-item">
		<div class="najpop-item-name">
			<a href="/model/opel/1284/opel-insignia-2017" title="2017 Opel Insignia (Insigni...">
				2017 Opel Insignia (Insigni...			</a>
		</div>
		<div class="najpop-item-img">
			<a href="/model/opel/1284/opel-insignia-2017" title="2017 Opel Insignia (Insigni...">
				<img src="/resources/images/variant/1938/thumb_insignia_5.jpg" 	title="Opel Insignia"
					 alt="Opel Insignia"
				 	 class="thumb-image">
			</a>
			<div class="ocena"><span>4.3</span>
			</div>
		</div>
		<div class="najpop-item-val">
			from: 			</br>
			<strong>12.300</strong>
			</br>EUR</div>
	</div>	
	<div class="najpop-item">
		<div class="najpop-item-name">
			<a href="/model/mercedes-benz/1378/mercedes-benz-c-2018" title="2018 Mercedes Benz C (W 205...">
				2018 Mercedes Benz C (W 205...			</a>
		</div>
		<div class="najpop-item-img">
			<a href="/model/mercedes-benz/1378/mercedes-benz-c-2018" title="2018 Mercedes Benz C (W 205...">
				<img src="/resources/images/variant/2086/thumb_2018_c_class_sedan.jpg" 	title="Mercedes Benz C"
					 alt="Mercedes Benz C"
				 	 class="thumb-image">
			</a>
			<div class="ocena"><span>4.8</span>
			</div>
		</div>
		<div class="najpop-item-val">
			from: 			</br>
			<strong>32.500</strong>
			</br>EUR</div>
	</div>		</div>
	</div>

			</div>
		<div class="side-ad">
	<script async src="//pagead2.googlesyndication.com/pagead/js/adsbygoogle.js"></script>
	<ins class="adsbygoogle side-ad-vertical"
	     style="width:160px;height:600px"
	     data-ad-client="ca-pub-2798966384669812"
	     data-ad-slot="9565981886"></ins>
	<script>
	     (adsbygoogle = window.adsbygoogle || []).push({});
	</script>

	<ins class="adsbygoogle side-ad-horizontal"
	     data-ad-client="ca-pub-2798966384669812"
	     data-ad-slot="8347642282"
	     data-full-width-responsive="true"></ins>
	<script>
	     (adsbygoogle = window.adsbygoogle || []).push({});
	</script>
</div>
	</div>
	<div id="page-wrap">
<div id="breadcrumb-wrap">
	<div class="breadcrumb-nav"><a href='/'>HOME page</a> / <a href='/specs'>specifications</a> / <a href='/make/audi'>Audi</a> / 2018 A4</div>
	<div id="sidebar-dodatno" class="search-box">
	<div class="sidebar-dodatno-wrap">
		<script async src="https://cse.google.com/cse.js?cx=016645478549947938819:d5ilzoxfr9e"></script>
		<script>
			window.onload = function(){
				$(".gsc-input").attr("placeholder", "Search this website...");
			}
		</script>
		<div class="gcse-search"></div>
	</div>
</div>	<div class="facebook">
				<div class="fb-like" data-href="https://www.automaniac.org/model/audi/1394/audi-a4-2018" data-width="" data-layout="button_count" data-action="like" data-size="small" data-show-faces="true" data-share="true"></div>
	</div>
</div>	<div id="modeli-auto-detaljnije">
		<div class="box-wrap">
			<div class="predlog-auto-logo">
				<a href="">
					<img src="/resources/images/make/audi.jpg">
				</a>
			</div>
			<div class="predlog-auto-naslov">
				Audi A4 (2018 - )				<div class="predlog-auto-stats">
					Audi B9 restyle, 					D - Segment (Large family car)				</div>
			</div>
			<div class="predlog-auto-slika">
								<div class="auto-score-wrap">
					<div class="score-left">
						<div class="score3b">
							<p>drivers'</p>
							<span>3.0</span>
							<p>rating</p>
						</div>
					</div>
					<div class="score-right clearfix">
						<div class="score5a"></div>
						<div class="score4a"></div>
						<div class="score3a"></div>
						<div class="score2a"></div>
						<div class="score1a"></div>
					</div>
				</div>
				<img src="/resources/images/model/1394/2018_audi_a4_interior.jpg"
					 id="main-model-image"/>
				<div class="auto-thumbs-wrap">
										<div class="auto-thumbs">
						<img src="/resources/images/model/1394/thumb_2018_audi_a4_interior.jpg"
							 alt="2018 Audi A4 (interior)"
							 title="2018 Audi A4 (interior)"
							 class="thumb-image" />
					</div>
											<div class="auto-thumbs">
							<img src="/resources/images/variant/2107/thumb_2018_audi_a4.jpg"
								 alt="2018 Audi A4 "
								 title="2018 Audi A4 "
								 class="thumb-image" />
						</div>
											<div class="auto-thumbs">
							<img src="/resources/images/variant/2108/thumb_2018_audi_a4_avant.jpg"
								 alt="2018 Audi A4 Avant"
								 title="2018 Audi A4 Avant"
								 class="thumb-image" />
						</div>
											<div class="auto-thumbs">
							<img src="/resources/images/variant/2109/thumb_2018_audi_s4.jpg"
								 alt="2018 Audi A4 S4"
								 title="2018 Audi A4 S4"
								 class="thumb-image" />
						</div>
											<div class="auto-thumbs">
							<img src="/resources/images/variant/2110/thumb_2018_audi_s4_avant.jpg"
								 alt="2018 Audi A4 S4 Avant"
								 title="2018 Audi A4 S4 Avant"
								 class="thumb-image" />
						</div>
											<div class="auto-thumbs">
							<img src="/resources/images/variant/2160/thumb_2018_audi_rs4_avant.jpg"
								 alt="2018 Audi A4 RS4 Avant"
								 title="2018 Audi A4 RS4 Avant"
								 class="thumb-image" />
						</div>
									</div>
			</div>
		</div>
					<div>
				<h2>
					<i class="fa fa-car" aria-hidden="true"></i>
					All the variants of Audi A4				</h2>
			</div>
				<div class="box-wrap">
			<div id="model-auto-wrap">
			<div class="model-auto-box variant-header accordion-controller bodystyle-header"
	 related-accordion="accordion-2107" variant_id="2107">
	<div class="model-auto-thumb">
		<img src="/resources/images/variant/2107/thumb_2018_audi_a4.jpg"
		     alt="A4 "
		     title="A4 "
		     class="thumb-image"/>
	</div>
	<div class="model-auto-i1">A4 </div>
	<div class="model-auto-i2">sedan - 4 door</div>
	<div class="model-auto-tip">
		<img src="/resources/images/category/sedan.png" />
	</div>
</div>
<div id="accordion-2107"
	 class="model-auto-content-wrap variant-accordion ">
			<div class="model-auto-podaci-wrap">
			<div class="model-auto-podaci-box">
				<div class="model-auto-puna-oznaka">
					<img class="fuel-badge" src="/resources/images/fuel/fuel1.png" />
					<a href="/audi/2018/14027/audi-a4-35-tfsi"
					   class="orange">
						Audi A4  35 TFSI					</a>
				</div>
			</div>
			<div class="model-auto-podaci-data">
				<div class="model-auto-gorivo">
					Fuel:<span>Petrol</span>
				</div>
				<div class="model-auto-pogon">
					Wheel drive:<span>Front</span>
				</div>
				<div class="model-auto-kubikaza">
											Displacement:<span>1984 cm3</span>
									</div>
				<div class="model-auto-snaga">
					Power:<span>150 ks</span>
				</div>
			</div>
		</div>
			<div class="model-auto-podaci-wrap">
			<div class="model-auto-podaci-box">
				<div class="model-auto-puna-oznaka">
					<img class="fuel-badge" src="/resources/images/fuel/fuel1.png" />
					<a href="/audi/2018/14028/audi-a4-40-tfsi"
					   class="orange">
						Audi A4  40 TFSI					</a>
				</div>
			</div>
			<div class="model-auto-podaci-data">
				<div class="model-auto-gorivo">
					Fuel:<span>Petrol</span>
				</div>
				<div class="model-auto-pogon">
					Wheel drive:<span>Front</span>
				</div>
				<div class="model-auto-kubikaza">
											Displacement:<span>1984 cm3</span>
									</div>
				<div class="model-auto-snaga">
					Power:<span>190 ks</span>
				</div>
			</div>
		</div>
			<div class="model-auto-podaci-wrap">
			<div class="model-auto-podaci-box">
				<div class="model-auto-puna-oznaka">
					<img class="fuel-badge" src="/resources/images/fuel/fuel1.png" />
					<a href="/audi/2018/14030/audi-a4-45-tfsi-quattro"
					   class="orange">
						Audi A4  45 TFSI quattro					</a>
				</div>
			</div>
			<div class="model-auto-podaci-data">
				<div class="model-auto-gorivo">
					Fuel:<span>Petrol</span>
				</div>
				<div class="model-auto-pogon">
					Wheel drive:<span>4 x 4</span>
				</div>
				<div class="model-auto-kubikaza">
											Displacement:<span>1984 cm3</span>
									</div>
				<div class="model-auto-snaga">
					Power:<span>245 ks</span>
				</div>
			</div>
		</div>
			<div class="model-auto-podaci-wrap">
			<div class="model-auto-podaci-box">
				<div class="model-auto-puna-oznaka">
					<img class="fuel-badge" src="/resources/images/fuel/fuel1.png" />
					<a href="/audi/2018/14029/audi-a4-45-tfsi"
					   class="orange">
						Audi A4  45 TFSI					</a>
				</div>
			</div>
			<div class="model-auto-podaci-data">
				<div class="model-auto-gorivo">
					Fuel:<span>Petrol</span>
				</div>
				<div class="model-auto-pogon">
					Wheel drive:<span>Front</span>
				</div>
				<div class="model-auto-kubikaza">
											Displacement:<span>1984 cm3</span>
									</div>
				<div class="model-auto-snaga">
					Power:<span>245 ks</span>
				</div>
			</div>
		</div>
			<div class="model-auto-podaci-wrap">
			<div class="model-auto-podaci-box">
				<div class="model-auto-puna-oznaka">
					<img class="fuel-badge" src="/resources/images/fuel/fuel2.png" />
					<a href="/audi/2018/14031/audi-a4-30-tdi"
					   class="orange">
						Audi A4  30 TDI					</a>
				</div>
			</div>
			<div class="model-auto-podaci-data">
				<div class="model-auto-gorivo">
					Fuel:<span>Diesel</span>
				</div>
				<div class="model-auto-pogon">
					Wheel drive:<span>Front</span>
				</div>
				<div class="model-auto-kubikaza">
											Displacement:<span>1968 cm3</span>
									</div>
				<div class="model-auto-snaga">
					Power:<span>122 ks</span>
				</div>
			</div>
		</div>
			<div class="model-auto-podaci-wrap">
			<div class="model-auto-podaci-box">
				<div class="model-auto-puna-oznaka">
					<img class="fuel-badge" src="/resources/images/fuel/fuel2.png" />
					<a href="/audi/2018/14032/audi-a4-35-tdi"
					   class="orange">
						Audi A4  35 TDI					</a>
				</div>
			</div>
			<div class="model-auto-podaci-data">
				<div class="model-auto-gorivo">
					Fuel:<span>Diesel</span>
				</div>
				<div class="model-auto-pogon">
					Wheel drive:<span>Front</span>
				</div>
				<div class="model-auto-kubikaza">
											Displacement:<span>1968 cm3</span>
									</div>
				<div class="model-auto-snaga">
					Power:<span>150 ks</span>
				</div>
			</div>
		</div>
			<div class="model-auto-podaci-wrap">
			<div class="model-auto-podaci-box">
				<div class="model-auto-puna-oznaka">
					<img class="fuel-badge" src="/resources/images/fuel/fuel2.png" />
					<a href="/audi/2018/14034/audi-a4-40-tdi-quattro"
					   class="orange">
						Audi A4  40 TDI quattro					</a>
				</div>
			</div>
			<div class="model-auto-podaci-data">
				<div class="model-auto-gorivo">
					Fuel:<span>Diesel</span>
				</div>
				<div class="model-auto-pogon">
					Wheel drive:<span>4 x 4</span>
				</div>
				<div class="model-auto-kubikaza">
											Displacement:<span>1968 cm3</span>
									</div>
				<div class="model-auto-snaga">
					Power:<span>190 ks</span>
				</div>
			</div>
		</div>
			<div class="model-auto-podaci-wrap">
			<div class="model-auto-podaci-box">
				<div class="model-auto-puna-oznaka">
					<img class="fuel-badge" src="/resources/images/fuel/fuel2.png" />
					<a href="/audi/2018/14033/audi-a4-40-tdi"
					   class="orange">
						Audi A4  40 TDI					</a>
				</div>
			</div>
			<div class="model-auto-podaci-data">
				<div class="model-auto-gorivo">
					Fuel:<span>Diesel</span>
				</div>
				<div class="model-auto-pogon">
					Wheel drive:<span>Front</span>
				</div>
				<div class="model-auto-kubikaza">
											Displacement:<span>1968 cm3</span>
									</div>
				<div class="model-auto-snaga">
					Power:<span>190 ks</span>
				</div>
			</div>
		</div>
			<div class="model-auto-podaci-wrap">
			<div class="model-auto-podaci-box">
				<div class="model-auto-puna-oznaka">
					<img class="fuel-badge" src="/resources/images/fuel/fuel2.png" />
					<a href="/audi/2018/14035/audi-a4-45-tdi-quattro"
					   class="orange">
						Audi A4  45 TDI quattro					</a>
				</div>
			</div>
			<div class="model-auto-podaci-data">
				<div class="model-auto-gorivo">
					Fuel:<span>Diesel</span>
				</div>
				<div class="model-auto-pogon">
					Wheel drive:<span>4 x 4</span>
				</div>
				<div class="model-auto-kubikaza">
											Displacement:<span>2967 cm3</span>
									</div>
				<div class="model-auto-snaga">
					Power:<span>231 ks</span>
				</div>
			</div>
		</div>
			<div class="model-auto-podaci-wrap">
			<div class="model-auto-podaci-box">
				<div class="model-auto-puna-oznaka">
					<img class="fuel-badge" src="/resources/images/fuel/fuel2.png" />
					<a href="/audi/2018/14036/audi-a4-50-tdi-quattro"
					   class="orange">
						Audi A4  50 TDI quattro					</a>
				</div>
			</div>
			<div class="model-auto-podaci-data">
				<div class="model-auto-gorivo">
					Fuel:<span>Diesel</span>
				</div>
				<div class="model-auto-pogon">
					Wheel drive:<span>4 x 4</span>
				</div>
				<div class="model-auto-kubikaza">
											Displacement:<span>2967 cm3</span>
									</div>
				<div class="model-auto-snaga">
					Power:<span>286 ks</span>
				</div>
			</div>
		</div>
	</div><div class="model-auto-box variant-header accordion-controller bodystyle-header"
	 related-accordion="accordion-2108" variant_id="2108">
	<div class="model-auto-thumb">
		<img src="/resources/images/variant/2108/thumb_2018_audi_a4_avant.jpg"
		     alt="A4 Avant"
		     title="A4 Avant"
		     class="thumb-image"/>
	</div>
	<div class="model-auto-i1">A4 Avant</div>
	<div class="model-auto-i2">wagon - 5 door</div>
	<div class="model-auto-tip">
		<img src="/resources/images/category/wagon.png" />
	</div>
</div>
<div id="accordion-2108"
	class="model-auto-content-wrap variant-accordion ">
			<div class="model-auto-podaci-wrap">
			<div class="model-auto-podaci-box">
				<div class="model-auto-puna-oznaka">
					<img class="fuel-badge" src="/resources/images/fuel/fuel1.png" />
					<a href="/audi/2018/14037/audi-a4-35-tfsi-avant"
	class="orange">
						Audi A4 Avant 35 TFSI					</a>
				</div>
			</div>
			<div class="model-auto-podaci-data">
				<div class="model-auto-gorivo">
					Fuel:<span>Petrol</span>
				</div>
				<div class="model-auto-pogon">
					Wheel drive:<span>Front</span>
				</div>
				<div class="model-auto-kubikaza">
											Displacement:<span>1984 cm3</span>
									</div>
				<div class="model-auto-snaga">
					Power:<span>150 ks</span>
				</div>
			</div>
		</div>
			<div class="model-auto-podaci-wrap">
			<div class="model-auto-podaci-box">
				<div class="model-auto-puna-oznaka">
					<img class="fuel-badge" src="/resources/images/fuel/fuel1.png" />
					<a href="/audi/2018/14038/audi-a4-40-tfsi-avant"
					   class="orange">
						Audi A4 Avant 40 TFSI					</a>
				</div>
			</div>
			<div class="model-auto-podaci-data">
				<div class="model-auto-gorivo">
					Fuel:<span>Petrol</span>
				</div>
				<div class="model-auto-pogon">
					Wheel drive:<span>Front</span>
				</div>
				<div class="model-auto-kubikaza">
											Displacement:<span>1984 cm3</span>
									</div>
				<div class="model-auto-snaga">
					Power:<span>190 ks</span>
				</div>
			</div>
		</div>
			<div class="model-auto-podaci-wrap">
			<div class="model-auto-podaci-box">
				<div class="model-auto-puna-oznaka">
					<img class="fuel-badge" src="/resources/images/fuel/fuel1.png" />
					<a href="/audi/2018/14040/audi-a4-45-tfsi-avant"
					   class="orange">
						Audi A4 Avant 45 TFSI					</a>
				</div>
			</div>
			<div class="model-auto-podaci-data">
				<div class="model-auto-gorivo">
					Fuel:<span>Petrol</span>
				</div>
				<div class="model-auto-pogon">
					Wheel drive:<span>Front</span>
				</div>
				<div class="model-auto-kubikaza">
											Displacement:<span>1984 cm3</span>
									</div>
				<div class="model-auto-snaga">
					Power:<span>245 ks</span>
				</div>
			</div>
		</div>
			<div class="model-auto-podaci-wrap">
			<div class="model-auto-podaci-box">
				<div class="model-auto-puna-oznaka">
					<img class="fuel-badge" src="/resources/images/fuel/fuel1.png" />
					<a href="/audi/2018/14039/audi-a4-45-tfsi-quattro-avant"
					   class="orange">
						Audi A4 Avant 45 TFSI quattro					</a>
				</div>
			</div>
			<div class="model-auto-podaci-data">
				<div class="model-auto-gorivo">
					Fuel:<span>Petrol</span>
				</div>
				<div class="model-auto-pogon">
					Wheel drive:<span>4 x 4</span>
				</div>
				<div class="model-auto-kubikaza">
											Displacement:<span>1984 cm3</span>
									</div>
				<div class="model-auto-snaga">
					Power:<span>245 ks</span>
				</div>
			</div>
		</div>
			<div class="model-auto-podaci-wrap">
			<div class="model-auto-podaci-box">
				<div class="model-auto-puna-oznaka">
					<img class="fuel-badge" src="/resources/images/fuel/fuel2.png" />
					<a href="/audi/2018/14041/audi-a4-30-tdi-avant"
					   class="orange">
						Audi A4 Avant 30 TDI					</a>
				</div>
			</div>
			<div class="model-auto-podaci-data">
				<div class="model-auto-gorivo">
					Fuel:<span>Diesel</span>
				</div>
				<div class="model-auto-pogon">
					Wheel drive:<span>Front</span>
				</div>
				<div class="model-auto-kubikaza">
											Displacement:<span>1968 cm3</span>
									</div>
				<div class="model-auto-snaga">
					Power:<span>122 ks</span>
				</div>
			</div>
		</div>
			<div class="model-auto-podaci-wrap">
			<div class="model-auto-podaci-box">
				<div class="model-auto-puna-oznaka">
					<img class="fuel-badge" src="/resources/images/fuel/fuel2.png" />
					<a href="/audi/2018/14042/audi-a4-35-tdi-avant"
					   class="orange">
						Audi A4 Avant 35 TDI					</a>
				</div>
			</div>
			<div class="model-auto-podaci-data">
				<div class="model-auto-gorivo">
					Fuel:<span>Diesel</span>
				</div>
				<div class="model-auto-pogon">
					Wheel drive:<span>Front</span>
				</div>
				<div class="model-auto-kubikaza">
											Displacement:<span>1968 cm3</span>
									</div>
				<div class="model-auto-snaga">
					Power:<span>150 ks</span>
				</div>
			</div>
		</div>
			<div class="model-auto-podaci-wrap">
			<div class="model-auto-podaci-box">
				<div class="model-auto-puna-oznaka">
					<img class="fuel-badge" src="/resources/images/fuel/fuel2.png" />
					<a href="/audi/2018/14044/audi-a4-40-tdi-avant"
					   class="orange">
						Audi A4 Avant 40 TDI					</a>
				</div>
			</div>
			<div class="model-auto-podaci-data">
				<div class="model-auto-gorivo">
					Fuel:<span>Diesel</span>
				</div>
				<div class="model-auto-pogon">
					Wheel drive:<span>Front</span>
				</div>
				<div class="model-auto-kubikaza">
											Displacement:<span>1968 cm3</span>
									</div>
				<div class="model-auto-snaga">
					Power:<span>190 ks</span>
				</div>
			</div>
		</div>
			<div class="model-auto-podaci-wrap">
			<div class="model-auto-podaci-box">
				<div class="model-auto-puna-oznaka">
					<img class="fuel-badge" src="/resources/images/fuel/fuel2.png" />
					<a href="/audi/2018/14043/audi-a4-40-tdi-quattro-avant"
					   class="orange">
						Audi A4 Avant 40 TDI quattro					</a>
				</div>
			</div>
			<div class="model-auto-podaci-data">
				<div class="model-auto-gorivo">
					Fuel:<span>Diesel</span>
				</div>
				<div class="model-auto-pogon">
					Wheel drive:<span>4 x 4</span>
				</div>
				<div class="model-auto-kubikaza">
											Displacement:<span>1968 cm3</span>
									</div>
				<div class="model-auto-snaga">
					Power:<span>190 ks</span>
				</div>
			</div>
		</div>
			<div class="model-auto-podaci-wrap">
			<div class="model-auto-podaci-box">
				<div class="model-auto-puna-oznaka">
					<img class="fuel-badge" src="/resources/images/fuel/fuel2.png" />
					<a href="/audi/2018/14045/audi-a4-45-tdi-quattro-avant"
					   class="orange">
						Audi A4 Avant 45 TDI quattro					</a>
				</div>
			</div>
			<div class="model-auto-podaci-data">
				<div class="model-auto-gorivo">
					Fuel:<span>Diesel</span>
				</div>
				<div class="model-auto-pogon">
					Wheel drive:<span>4 x 4</span>
				</div>
				<div class="model-auto-kubikaza">
											Displacement:<span>2967 cm3</span>
									</div>
				<div class="model-auto-snaga">
					Power:<span>231 ks</span>
				</div>
			</div>
		</div>
			<div class="model-auto-podaci-wrap">
			<div class="model-auto-podaci-box">
				<div class="model-auto-puna-oznaka">
					<img class="fuel-badge" src="/resources/images/fuel/fuel2.png" />
					<a href="/audi/2018/14046/audi-a4-50-tdi-quattro-avant"
					   class="orange">
						Audi A4 Avant 50 TDI quattro					</a>
				</div>
			</div>
			<div class="model-auto-podaci-data">
				<div class="model-auto-gorivo">
					Fuel:<span>Diesel</span>
				</div>
				<div class="model-auto-pogon">
					Wheel drive:<span>4 x 4</span>
				</div>
				<div class="model-auto-kubikaza">
											Displacement:<span>2967 cm3</span>
									</div>
				<div class="model-auto-snaga">
					Power:<span>286 ks</span>
				</div>
			</div>
		</div>
	</div><div class="model-auto-box variant-header accordion-controller bodystyle-header"
	 related-accordion="accordion-2109" variant_id="2109">
	<div class="model-auto-thumb">
		<img src="/resources/images/variant/2109/thumb_2018_audi_s4.jpg"
		     alt="A4 S4"
		     title="A4 S4"
		     class="thumb-image"/>
	</div>
	<div class="model-auto-i1">A4 S4</div>
	<div class="model-auto-i2">sedan - 4 door</div>
	<div class="model-auto-tip">
		<img src="/resources/images/category/sedan.png" />
	</div>
</div>
<div id="accordion-2109"
	 class="model-auto-content-wrap variant-accordion first-accordion">
			<div class="model-auto-podaci-wrap">
			<div class="model-auto-podaci-box">
				<div class="model-auto-puna-oznaka">
					<img class="fuel-badge" src="/resources/images/fuel/fuel2.png" />
					<a href="/audi/2018/14047/audi-a4-tdi-quattro-s4"
					   class="orange">
						Audi A4 S4 TDI quattro					</a>
				</div>
			</div>
			<div class="model-auto-podaci-data">
				<div class="model-auto-gorivo">
					Fuel:<span>Diesel</span>
				</div>
				<div class="model-auto-pogon">
					Wheel drive:<span>4 x 4</span>
				</div>
				<div class="model-auto-kubikaza">
											Displacement:<span>2967 cm3</span>
									</div>
				<div class="model-auto-snaga">
					Power:<span>347 ks</span>
				</div>
			</div>
		</div>
	</div><div class="model-auto-box variant-header accordion-controller bodystyle-header"
	 related-accordion="accordion-2110" variant_id="2110">
	<div class="model-auto-thumb">
		<img src="/resources/images/variant/2110/thumb_2018_audi_s4_avant.jpg"
		     alt="A4 S4 Avant"
		     title="A4 S4 Avant"
		     class="thumb-image"/>
	</div>
	<div class="model-auto-i1">A4 S4 Avant</div>
	<div class="model-auto-i2">wagon - 5 door</div>
	<div class="model-auto-tip">
		<img src="/resources/images/category/wagon.png" />
	</div>
</div>
<div id="accordion-2110"
	 class="model-auto-content-wrap variant-accordion ">
			<div class="model-auto-podaci-wrap">
			<div class="model-auto-podaci-box">
				<div class="model-auto-puna-oznaka">
					<img class="fuel-badge" src="/resources/images/fuel/fuel2.png" />
					<a href="/audi/2018/14048/audi-a4-tdi-quattro-s4-avant"
					   class="orange">
						Audi A4 S4 Avant TDI quattro					</a>
				</div>
			</div>
			<div class="model-auto-podaci-data">
				<div class="model-auto-gorivo">
					Fuel:<span>Diesel</span>
				</div>
				<div class="model-auto-pogon">
					Wheel drive:<span>4 x 4</span>
				</div>
				<div class="model-auto-kubikaza">
											Displacement:<span>2967 cm3</span>
									</div>
				<div class="model-auto-snaga">
					Power:<span>347 ks</span>
				</div>
			</div>
		</div>
	</div><div class="model-auto-box variant-header accordion-controller bodystyle-header"
	 related-accordion="accordion-2160" variant_id="2160">
	<div class="model-auto-thumb">
		<img src="/resources/images/variant/2160/thumb_2018_audi_rs4_avant.jpg"
		     alt="A4 RS4 Avant"
		     title="A4 RS4 Avant"
		     class="thumb-image"/>
	</div>
	<div class="model-auto-i1">A4 RS4 Avant</div>
	<div class="model-auto-i2">wagon - 5 door</div>
	<div class="model-auto-tip">
		<img src="/resources/images/category/wagon.png" />
	</div>
</div>
<div id="accordion-2160"
	 class="model-auto-content-wrap variant-accordion ">
			<div class="model-auto-podaci-wrap">
			<div class="model-auto-podaci-box">
				<div class="model-auto-puna-oznaka">
					<img class="fuel-badge" src="/resources/images/fuel/fuel1.png" />
					<a href="/audi/2018/14401/audi-a4-2.9-tfsi-quattro-rs4-avant"
					   class="orange">
						Audi A4 RS4 Avant 2.9 TFSI Quattro					</a>
				</div>
			</div>
			<div class="model-auto-podaci-data">
				<div class="model-auto-gorivo">
					Fuel:<span>Petrol</span>
				</div>
				<div class="model-auto-pogon">
					Wheel drive:<span>4 x 4</span>
				</div>
				<div class="model-auto-kubikaza">
											Displacement:<span>2894 cm3</span>
									</div>
				<div class="model-auto-snaga">
					Power:<span>450 ks</span>
				</div>
			</div>
		</div>
	</div>			</div>
		</div>
			</div>
	<div class="google-add-small">
		
	</div>
<div>
	<h2>
		<i class="fa fa-shield" aria-hidden="true"></i>
		Safety results for Audi A4	</h2>
</div>
<div class="box-wrap">
	<div class="euro-ncap-youtube">
		<div class="ncap-img">
			<img src="/resources/images/gui/details/euroncap.jpg">
		</div>
		<div class="stars">
			<div class="stars5"></div>
		</div>
		<div class="euro-ncap strong">
			EuroNCAP:
			90% occupant safety			
		</div>
					<iframe width="100%"
					height="360"
					src="https://www.youtube.com/embed/I37LkwLc2Qw?rel=0"
					frameborder="0"
					allow="accelerometer; autoplay; encrypted-media; gyroscope; picture-in-picture"
					allowfullscreen>
			</iframe>
			</div>
</div><div>
	<h2><i class="fa fa-search" aria-hidden="true"></i>Check a car by its VIN number</h2>
</div>
<div class="box-wrap" style="height: 250px;">
	<div
	  data-cvaff
	  data-locale="en-EU"
	  data-a="automanijak"
	  data-b="81ec5429"
	  style="width: 100%; height: 100%">
	</div>
</div><div>
	<h2><i class="fa fa-user" aria-hidden="true"></i>Visitor comments</h2>
</div>
<div class="box-wrap">
	<div class="fb-comments" data-href="https://www.automaniac.orghttps://www.automaniac.org/model/audi/1394/audi-a4-2018" data-width="670" data-numposts="5"></div>
</div>
	</div>
			</div>
		</div>
		<div id="footer">
			<div class="footer-wrap">
				Copyright &copy; 2015 - 2022 automaniac.org - All rights reserved.
				</br>
				Powered by <strong>Gama1 Solutions</strong>
				<div class="footer-links">
					<div class="footer-menu">
						<ul>
							<li><a href="/">HOME page</a></li>
							<li><a href="/about">about</a></li>
							<li><a href="/terms">terms & conditions</a></li>
							<li><a href="/legal">legal note</a></li>
							<li><a href="/faq">frequent questions (FAQ)</a></li>
							<li><a href="/contact">contact us</a></li>
						</ul>
					</div>

					<div class='footer-auto-brend'><ul><li><a href='/make/alfa-romeo'>Alfa Romeo</a></li><li><a href='/make/audi'>Audi</a></li><li><a href='/make/bmw'>BMW</a></li><li><a href='/make/chery'>Chery</a></li><li><a href='/make/chevrolet'>Chevrolet</a></li><li><a href='/make/chrysler'>Chrysler</a></li></ul></div><div class='footer-auto-brend'><ul><li><a href='/make/citroen'>Citroen</a></li><li><a href='/make/dacia'>Dacia</a></li><li><a href='/make/daihatsu'>Daihatsu</a></li><li><a href='/make/dodge'>Dodge</a></li><li><a href='/make/fiat'>FIAT</a></li><li><a href='/make/ford'>Ford</a></li></ul></div><div class='footer-auto-brend'><ul><li><a href='/make/honda'>Honda</a></li><li><a href='/make/hyundai'>Hyundai</a></li><li><a href='/make/infiniti'>Infiniti</a></li><li><a href='/make/jaguar'>Jaguar</a></li><li><a href='/make/jeep'>Jeep</a></li><li><a href='/make/kia'>KIA</a></li></ul></div><div class='footer-auto-brend'><ul><li><a href='/make/lada'>Lada</a></li><li><a href='/make/lancia'>Lancia</a></li><li><a href='/make/land-rover'>Land Rover</a></li><li><a href='/make/lexus'>Lexus</a></li><li><a href='/make/mazda'>Mazda</a></li><li><a href='/make/mercedes-benz'>Mercedes</a></li></ul></div><div class='footer-auto-brend'><ul><li><a href='/make/mini'>Mini</a></li><li><a href='/make/mitsubishi'>Mitsubishi</a></li><li><a href='/make/nissan'>Nissan</a></li><li><a href='/make/opel'>Opel</a></li><li><a href='/make/peugeot'>Peugeot</a></li><li><a href='/make/porsche'>Porsche</a></li></ul></div><div class='footer-auto-brend'><ul><li><a href='/make/renault'>Renault</a></li><li><a href='/make/rover'>Rover</a></li><li><a href='/make/saab'>SAAB</a></li><li><a href='/make/ssangyong'>SSangYong</a></li><li><a href='/make/seat'>Seat</a></li><li><a href='/make/smart'>Smart</a></li></ul></div><div class='footer-auto-brend'><ul><li><a href='/make/subaru'>Subaru</a></li><li><a href='/make/suzuki'>Suzuki</a></li><li><a href='/make/tesla'>Tesla</a></li><li><a href='/make/toyota'>Toyota</a></li><li><a href='/make/volkswagen'>Volkswagen</a></li><li><a href='/make/volvo'>Volvo</a></li></ul></div><div class='footer-auto-brend'><ul><li><a href='/make/zastava'>Zastava</a></li><li><a href='/make/škoda'>Škoda</a></li></ul></div>
				</div>
			</div>
		</div>
	</body>
</html>`
const model_html = `

<!DOCTYPE html>
<html lang="en">
	<head>
		<meta charset="utf-8">
		<title>
			Audi A1 (2010 - 2014) - AutoManiac		</title>
		<meta content="width=device-width, initial-scale=1.0" name="viewport">
		<meta content="Audi A1 specifications,Audi A1 consumption,Audi A1 crash test,Audi A1 prices,Audi A1 buying,Audi A1 selling,Audi A1 opinions" name="keywords">
		<meta content="2010 Audi A1 - Information, specifications and drivers opinions" name="description">
		<meta http-equiv="Cache-Control" content="no-cache, max-age=0, must-revalidate, no-store">
		<meta property="og:site_name" content="AutoManiac"/>
		<meta property="og:image" content="/resources/images/model/212/2010_audi_a1_interior.jpg" />
		<meta property="og:image:width" content="700"/>
		<meta property="og:image:height" content="300"/>
		<meta property="og:description" content="2010 Audi A1 - Information, specifications and drivers opinions" />
		<meta property="og:title" content="Audi A1 (2010 - 2014)" />
		<meta property="fb:app_id" content="290604621342699" />
		<meta property="fb:admins" content="1276696339"/>
		<link rel="image_src" type="image/jpeg" href="/resources/images/model/212/2010_audi_a1_interior.jpg" />
		<link rel="shortcut icon" href="/resources/images/gui/favicon.png" type="image/x-icon">
		<link rel="icon" href="/resources/images/gui/favicon.png" type="image/x-icon">
			<script type="text/javascript" src="/resources/js/jquery-1.8.3.js"></script>
	<script type="text/javascript" src="/resources/js/jquery-ui.js"></script>
	<script type="text/javascript" src="/resources/js/scripts.js?version=v2.1.0"></script>
	<script type="text/javascript" src="/resources/js/jquery.chosen.min.js"></script>
	<script src="https://apis.google.com/js/platform.js" async defer></script>
	<script async src="https://securepubads.g.doubleclick.net/tag/js/gpt.js"></script>
	<!--
	<script async src="//pagead2.googlesyndication.com/pagead/js/adsbygoogle.js"></script>
	<script type="text/javascript" src="/resources/js/adsense.js"></script>
	-->
	<script async defer crossorigin="anonymous" src="https://connect.facebook.net/en_GB/sdk.js#xfbml=1&version=v4.0"></script>
	
						
	<link rel="stylesheet" href="/resources/css/styles.css?version=v2.1.6" type="text/css" media="screen" charset="utf-8" />
	<link rel="stylesheet" href="/resources/css/chosen.css" type="text/css" media="screen" charset="utf-8" />
	<link rel="stylesheet" href="/resources/css/font-awesome.min.css" >
	<link rel="stylesheet" href="https://fonts.googleapis.com/css?family=Open+Sans:300,300i,400,400i,700,700i|Raleway:300,400,500,700,800|Montserrat:300,400,700">
		<link rel="icon" href="/resources/images/gui/favicon.png" >
		<link rel="shortcut icon" href="/resources/images/gui/favicon.png" type="image/x-icon">
		<link rel="apple-touch-icon" href="images/gui/apple-touch-icon.png">
	</head>

	<body root-url="">

		<!-- Load Facebook SDK for JavaScript -->
		<div id="fb-root"></div>
		<script>
		  window.fbAsyncInit = function() {
		    FB.init({
		      appId      : '290604621342699',
		      xfbml      : true,
		      version    : 'v2.8'
		    });
		  };

		  (function(d, s, id){
		     var js, fjs = d.getElementsByTagName(s)[0];
		     if (d.getElementById(id)) {return;}
		     js = d.createElement(s); js.id = id;
		     js.src = "//connect.facebook.net/en_US/sdk.js";
		     fjs.parentNode.insertBefore(js, fjs);
		   }(document, 'script', 'facebook-jssdk'));

		  (function(w,d,u,h,s){
		    h=d.getElementsByTagName('head')[0];
		    s=d.createElement('script');
		    s.async=1;
		    s.src=u+'/sdk.js';
		    h.appendChild(s);
		  })(window,document,'https://aff.carvertical.com');

			
			  (function(i,s,o,g,r,a,m){i['GoogleAnalyticsObject']=r;i[r]=i[r]||function(){
			  (i[r].q=i[r].q||[]).push(arguments)},i[r].l=1*new Date();a=s.createElement(o),
			  m=s.getElementsByTagName(o)[0];a.async=1;a.src=g;m.parentNode.insertBefore(a,m)
			  })(window,document,'script','//www.google-analytics.com/analytics.js','ga');
			
			  ga('create', 'UA-75322623-2', 'auto');
	  		  ga('send', 'pageview');

	  		
		
		</script>

		<div class="menu-wrap">
			<div class="topnav" id="myTopnav">
				<div class="top-bar">
					<div class="top-bar-wrap">
						<div class="top-bar-left">
							<a href="/about" title="about">about</a>
							<a href="/legal" title="legal note">legal note</a>
							<a href="/contact" title="contact us">contact us</a>
							<span style="color: #fff; font-weight: bold;"></span>
						</div>
						<div class="top-bar-right">
							<a href="#" title="">
								Login <i class="fa fa-user" aria-hidden="true"></i>
							</a>
							<a href="https://www.automanie.org" title="">
								deutsch (D, AT, CH)<i class="fa fa-globe" aria-hidden="true"></i>
							</a>
							<a href="https://www.automanijak.com" title="">
								srpski (ex-yugoslavia)<i class="fa fa-flag" aria-hidden="true"></i>
							</a>
						</div>
					</div>
				</div>
				<a href="#" class="icon"><i class="fa fa-bars"></i></a>
				<div class="bottom-bar">
					<div class="bottom-bar-wrap">
						<div class="bottom-bar-left">
							<div class="logo">
								<a href="/" title="">
									<img src="/resources/images/gui/logo.png" alt="" title="">
								</a>
							</div>
						</div>
						<div class="bottom-bar-right">
							<a href="/virtual-adviser">Virtual adviser</a>
							<a href="/specs">Specifications by model</a>
							<a href="/compare">Compare two cars</a>
							<a href="/badges">Automotive badges</a>
							<a href="/blog">Blog</a>
						</div>
					</div>
				</div>
			</div>
		</div>
		<div id="main-wrapper">
			<div class="left-ad">
				<script async src="//pagead2.googlesyndication.com/pagead/js/adsbygoogle.js"></script>
<!-- branding -->
<ins class="adsbygoogle branding-ad"
     style="width:300px;height:600px"
     data-ad-client="ca-pub-2798966384669812"
     data-ad-slot="1676022217"></ins>
<script>
(adsbygoogle = window.adsbygoogle || []).push({});
</script>				<script async src="https://pagead2.googlesyndication.com/pagead/js/adsbygoogle.js"></script>
<!-- massive-header -->
<ins class="adsbygoogle massive-header-ad"
     data-ad-client="ca-pub-2798966384669812"
     data-ad-slot="9878901748"
     data-full-width-responsive="true"></ins>
<script>
     (adsbygoogle = window.adsbygoogle || []).push({});
</script>			</div>
			<div class="right-ad">
				<script async src="//pagead2.googlesyndication.com/pagead/js/adsbygoogle.js"></script>
<!-- branding -->
<ins class="adsbygoogle branding-ad"
     style="width:300px;height:600px"
     data-ad-client="ca-pub-2798966384669812"
     data-ad-slot="1676022217"></ins>
<script>
(adsbygoogle = window.adsbygoogle || []).push({});
</script>			</div>
			<div id="wrapper" class="clearfix">	<div id="sidebar">
		<div class="sidebar-wrap">
		<div id="sidebar-dodatno" class="search-box">
	<div class="sidebar-dodatno-wrap">
		<script async src="https://cse.google.com/cse.js?cx=016645478549947938819:d5ilzoxfr9e"></script>
		<script>
			window.onload = function(){
				$(".gsc-input").attr("placeholder", "Search this website...");
			}
		</script>
		<div class="gcse-search"></div>
	</div>
</div>	
	<div id="sidebar-dodatno">
		<div class="sidebar-misljenje-naslov">
			<h2>Drivers' opinion</h2>
		</div>
		<div class="sidebar-dodatno-wrap">
		<div class="misljenje-vozaca">
	<span>Reliability <strong title="- Malfunction frequency&#13;- Malfunction severity&#13;- Spare parts availability">(?)</strong>:</span>
	<div class="stars">
		<div class="stars4"></div>
	</div>
	<div class="misljenje-ocena">4.0</div>
</div><div class="misljenje-vozaca">
	<span>Performance <strong title="- Acceleration&#13;- Top speed&#13;- Braking efficiency&#13;- Handling">(?)</strong>:</span>
	<div class="stars">
		<div class="stars4"></div>
	</div>
	<div class="misljenje-ocena">4.0</div>
</div><div class="misljenje-vozaca">
	<span>Economy <strong title="- Fuel consumption&#13;- Regular maintenance prices&#13;- Spare parts prices&#13;- Other expenses specific to this car">(?)</strong>:</span>
	<div class="stars">
		<div class="stars3"></div>
	</div>
	<div class="misljenje-ocena">3.0</div>
</div><div class="misljenje-vozaca">
	<span>Comfort <strong title="- Suspension quality&#13;- Seats comfort&#13;- Legs and head space&#13;- Driver's position & ergonomy">(?)</strong>:</span>
	<div class="stars">
		<div class="stars4"></div>
	</div>
	<div class="misljenje-ocena">3.5</div>
</div><div class="misljenje-vozaca">
	<span>Practicality <strong title="- Easy to use&#13;- Quantity & quality of in-car compartments&#13;- Luggage space access&#13;- Spare parts replacement ease">(?)</strong>:</span>
	<div class="stars">
		<div class="stars4"></div>
	</div>
	<div class="misljenje-ocena">3.5</div>
</div>		</div>
	</div><div id="sidebar-dodatno">
	<div class="sidebar-minimalna-cena-naslov">
		<h2>Starting price</h2>
	</div>
	<div class="sidebar-dodatno-wrap">
		<div class="dodatno-text">
			<div>
				from				</br>
				<strong>6.000</strong>
				</br>
				EUR*
			</div>
			<span>Starting price is informative only and is being updated using information from leading online yellow pages. Prices of well maintained examples are commonly 20% higher.</span> 
		</div>
	</div>
</div>	
	<div class="najpopulariji-wrap-bottom">
		<div class="sidebar-najpopularniji-naslov">
			<h2>Some similar models...</h2>
		</div>
		<div class="najpop-item-wrap">
			
	<div class="najpop-item">
		<div class="najpop-item-name">
			<a href="/model/smart/832/smart-forfour-2014" title="2014 Smart ForFour ">
				2014 Smart ForFour 			</a>
		</div>
		<div class="najpop-item-img">
			<a href="/model/smart/832/smart-forfour-2014" title="2014 Smart ForFour ">
				<img src="/resources/images/variant/1762/thumb_for_four_2.jpg" 	title="Smart ForFour"
					 alt="Smart ForFour"
				 	 class="thumb-image">
			</a>
			<div class="ocena"><span>3.0</span>
			</div>
		</div>
		<div class="najpop-item-val">
			from: 			</br>
			<strong>6.290</strong>
			</br>EUR</div>
	</div>	
	<div class="najpop-item">
		<div class="najpop-item-name">
			<a href="/model/volvo/995/volvo-c30-2009" title="2009 Volvo C30 ">
				2009 Volvo C30 			</a>
		</div>
		<div class="najpop-item-img">
			<a href="/model/volvo/995/volvo-c30-2009" title="2009 Volvo C30 ">
				<img src="/resources/images/variant/1612/thumb_c30_2.jpg" 	title="Volvo C30"
					 alt="Volvo C30"
				 	 class="thumb-image">
			</a>
			<div class="ocena"><span>4.4</span>
			</div>
		</div>
		<div class="najpop-item-val">
			from: 			</br>
			<strong>3.600</strong>
			</br>EUR</div>
	</div>	
	<div class="najpop-item">
		<div class="najpop-item-name">
			<a href="/model/citroen/1074/citroen-ds3-2014" title="2014 Citroen DS3 ">
				2014 Citroen DS3 			</a>
		</div>
		<div class="najpop-item-img">
			<a href="/model/citroen/1074/citroen-ds3-2014" title="2014 Citroen DS3 ">
				<img src="/resources/images/variant/371/thumb_ds3_3.jpg" 	title="Citroen DS3"
					 alt="Citroen DS3"
				 	 class="thumb-image">
			</a>
			<div class="ocena"><span>5.0</span>
			</div>
		</div>
		<div class="najpop-item-val">
			from: 			</br>
			<strong>6.700</strong>
			</br>EUR</div>
	</div>		</div>
	</div>

			</div>
		<div class="side-ad">
	<script async src="//pagead2.googlesyndication.com/pagead/js/adsbygoogle.js"></script>
	<ins class="adsbygoogle side-ad-vertical"
	     style="width:160px;height:600px"
	     data-ad-client="ca-pub-2798966384669812"
	     data-ad-slot="9565981886"></ins>
	<script>
	     (adsbygoogle = window.adsbygoogle || []).push({});
	</script>

	<ins class="adsbygoogle side-ad-horizontal"
	     data-ad-client="ca-pub-2798966384669812"
	     data-ad-slot="8347642282"
	     data-full-width-responsive="true"></ins>
	<script>
	     (adsbygoogle = window.adsbygoogle || []).push({});
	</script>
</div>
	</div>
	<div id="page-wrap">
<div id="breadcrumb-wrap">
	<div class="breadcrumb-nav"><a href='/'>HOME page</a> / <a href='/specs'>specifications</a> / <a href='/make/audi'>Audi</a> / 2010 A1</div>
	<div id="sidebar-dodatno" class="search-box">
	<div class="sidebar-dodatno-wrap">
		<script async src="https://cse.google.com/cse.js?cx=016645478549947938819:d5ilzoxfr9e"></script>
		<script>
			window.onload = function(){
				$(".gsc-input").attr("placeholder", "Search this website...");
			}
		</script>
		<div class="gcse-search"></div>
	</div>
</div>	<div class="facebook">
				<div class="fb-like" data-href="https://www.automaniac.org/model/audi/212/audi-a1-2010" data-width="" data-layout="button_count" data-action="like" data-size="small" data-show-faces="true" data-share="true"></div>
	</div>
</div>	<div id="modeli-auto-detaljnije">
		<div class="box-wrap">
			<div class="predlog-auto-logo">
				<a href="">
					<img src="/resources/images/make/audi.jpg">
				</a>
			</div>
			<div class="predlog-auto-naslov">
				Audi A1 (2010 - 2014)				<div class="predlog-auto-stats">
					Audi 8X, 					B - Segment (City car)				</div>
			</div>
			<div class="predlog-auto-slika">
								<div class="auto-score-wrap">
					<div class="score-left">
						<div class="score4b">
							<p>drivers'</p>
							<span>3.6</span>
							<p>rating</p>
						</div>
					</div>
					<div class="score-right clearfix">
						<div class="score5a"></div>
						<div class="score4a"></div>
						<div class="score3a"></div>
						<div class="score2a"></div>
						<div class="score1a"></div>
					</div>
				</div>
				<img src="/resources/images/model/212/2010_audi_a1_interior.jpg"
					 id="main-model-image"/>
				<div class="auto-thumbs-wrap">
										<div class="auto-thumbs">
						<img src="/resources/images/model/212/thumb_2010_audi_a1_interior.jpg"
							 alt="2010 Audi A1 (interior)"
							 title="2010 Audi A1 (interior)"
							 class="thumb-image" />
					</div>
											<div class="auto-thumbs">
							<img src="/resources/images/variant/22/thumb_a1_2.jpg"
								 alt="2010 Audi A1 "
								 title="2010 Audi A1 "
								 class="thumb-image" />
						</div>
											<div class="auto-thumbs">
							<img src="/resources/images/variant/23/thumb_a1_1.jpg"
								 alt="2010 Audi A1 Sportback"
								 title="2010 Audi A1 Sportback"
								 class="thumb-image" />
						</div>
									</div>
			</div>
		</div>
					<div>
				<h2>
					<i class="fa fa-car" aria-hidden="true"></i>
					All the variants of Audi A1				</h2>
			</div>
				<div class="box-wrap">
			<div id="model-auto-wrap">
			<div class="model-auto-box variant-header accordion-controller bodystyle-header"
	 related-accordion="accordion-22" variant_id="22">
	<div class="model-auto-thumb">
		<img src="/resources/images/variant/22/thumb_a1_2.jpg"
		     alt="A1 "
		     title="A1 "
		     class="thumb-image"/>
	</div>
	<div class="model-auto-i1">A1 </div>
	<div class="model-auto-i2">hatchback - 3 door</div>
	<div class="model-auto-tip">
		<img src="/resources/images/category/hatchback.png" />
	</div>
</div>
<div id="accordion-22"
	 class="model-auto-content-wrap variant-accordion first-accordion">
			<div class="model-auto-podaci-wrap">
			<div class="model-auto-podaci-box">
				<div class="model-auto-puna-oznaka">
					<img class="fuel-badge" src="/resources/images/fuel/fuel1.png" />
					<a href="/audi/2010/345/audi-a1-1.2-tfsi"
					   class="orange">
						Audi A1  1.2 TFSI					</a>
				</div>
			</div>
			<div class="model-auto-podaci-data">
				<div class="model-auto-gorivo">
					Fuel:<span>Petrol</span>
				</div>
				<div class="model-auto-pogon">
					Wheel drive:<span>Front</span>
				</div>
				<div class="model-auto-kubikaza">
											Displacement:<span>1197 cm3</span>
									</div>
				<div class="model-auto-snaga">
					Power:<span>86 ks</span>
				</div>
			</div>
		</div>
			<div class="model-auto-podaci-wrap">
			<div class="model-auto-podaci-box">
				<div class="model-auto-puna-oznaka">
					<img class="fuel-badge" src="/resources/images/fuel/fuel1.png" />
					<a href="/audi/2010/349/audi-a1-1.4-tfsi"
					   class="orange">
						Audi A1  1.4 TFSI					</a>
				</div>
			</div>
			<div class="model-auto-podaci-data">
				<div class="model-auto-gorivo">
					Fuel:<span>Petrol</span>
				</div>
				<div class="model-auto-pogon">
					Wheel drive:<span>Front</span>
				</div>
				<div class="model-auto-kubikaza">
											Displacement:<span>1390 cm3</span>
									</div>
				<div class="model-auto-snaga">
					Power:<span>122 ks</span>
				</div>
			</div>
		</div>
			<div class="model-auto-podaci-wrap">
			<div class="model-auto-podaci-box">
				<div class="model-auto-puna-oznaka">
					<img class="fuel-badge" src="/resources/images/fuel/fuel1.png" />
					<a href="/audi/2010/347/audi-a1-1.4-tfsi-cod"
					   class="orange">
						Audi A1  1.4 TFSI COD					</a>
				</div>
			</div>
			<div class="model-auto-podaci-data">
				<div class="model-auto-gorivo">
					Fuel:<span>Petrol</span>
				</div>
				<div class="model-auto-pogon">
					Wheel drive:<span>Front</span>
				</div>
				<div class="model-auto-kubikaza">
											Displacement:<span>1390 cm3</span>
									</div>
				<div class="model-auto-snaga">
					Power:<span>140 ks</span>
				</div>
			</div>
		</div>
			<div class="model-auto-podaci-wrap">
			<div class="model-auto-podaci-box">
				<div class="model-auto-puna-oznaka">
					<img class="fuel-badge" src="/resources/images/fuel/fuel1.png" />
					<a href="/audi/2010/348/audi-a1-1.4-tfsi-185"
					   class="orange">
						Audi A1  1.4 TFSI 185					</a>
				</div>
			</div>
			<div class="model-auto-podaci-data">
				<div class="model-auto-gorivo">
					Fuel:<span>Petrol</span>
				</div>
				<div class="model-auto-pogon">
					Wheel drive:<span>Front</span>
				</div>
				<div class="model-auto-kubikaza">
											Displacement:<span>1390 cm3</span>
									</div>
				<div class="model-auto-snaga">
					Power:<span>185 ks</span>
				</div>
			</div>
		</div>
			<div class="model-auto-podaci-wrap">
			<div class="model-auto-podaci-box">
				<div class="model-auto-puna-oznaka">
					<img class="fuel-badge" src="/resources/images/fuel/fuel2.png" />
					<a href="/audi/2010/350/audi-a1-1.6-tdi"
					   class="orange">
						Audi A1  1.6 TDI					</a>
				</div>
			</div>
			<div class="model-auto-podaci-data">
				<div class="model-auto-gorivo">
					Fuel:<span>Diesel</span>
				</div>
				<div class="model-auto-pogon">
					Wheel drive:<span>Front</span>
				</div>
				<div class="model-auto-kubikaza">
											Displacement:<span>1598 cm3</span>
									</div>
				<div class="model-auto-snaga">
					Power:<span>90 ks</span>
				</div>
			</div>
		</div>
			<div class="model-auto-podaci-wrap">
			<div class="model-auto-podaci-box">
				<div class="model-auto-puna-oznaka">
					<img class="fuel-badge" src="/resources/images/fuel/fuel2.png" />
					<a href="/audi/2010/352/audi-a1-1.6-tdi"
					   class="orange">
						Audi A1  1.6 TDI					</a>
				</div>
			</div>
			<div class="model-auto-podaci-data">
				<div class="model-auto-gorivo">
					Fuel:<span>Diesel</span>
				</div>
				<div class="model-auto-pogon">
					Wheel drive:<span>Front</span>
				</div>
				<div class="model-auto-kubikaza">
											Displacement:<span>1598 cm3</span>
									</div>
				<div class="model-auto-snaga">
					Power:<span>105 ks</span>
				</div>
			</div>
		</div>
			<div class="model-auto-podaci-wrap">
			<div class="model-auto-podaci-box">
				<div class="model-auto-puna-oznaka">
					<img class="fuel-badge" src="/resources/images/fuel/fuel2.png" />
					<a href="/audi/2010/353/audi-a1-2.0-tdi"
					   class="orange">
						Audi A1  2.0 TDI					</a>
				</div>
			</div>
			<div class="model-auto-podaci-data">
				<div class="model-auto-gorivo">
					Fuel:<span>Diesel</span>
				</div>
				<div class="model-auto-pogon">
					Wheel drive:<span>Front</span>
				</div>
				<div class="model-auto-kubikaza">
											Displacement:<span>1968 cm3</span>
									</div>
				<div class="model-auto-snaga">
					Power:<span>143 ks</span>
				</div>
			</div>
		</div>
	</div><div class="model-auto-box variant-header accordion-controller bodystyle-header"
	 related-accordion="accordion-23" variant_id="23">
	<div class="model-auto-thumb">
		<img src="/resources/images/variant/23/thumb_a1_1.jpg"
		     alt="A1 Sportback"
		     title="A1 Sportback"
		     class="thumb-image"/>
	</div>
	<div class="model-auto-i1">A1 Sportback</div>
	<div class="model-auto-i2">hatchback - 5 door</div>
	<div class="model-auto-tip">
		<img src="/resources/images/category/hatchback.png" />
	</div>
</div>
<div id="accordion-23"
	 class="model-auto-content-wrap variant-accordion ">
			<div class="model-auto-podaci-wrap">
			<div class="model-auto-podaci-box">
				<div class="model-auto-puna-oznaka">
					<img class="fuel-badge" src="/resources/images/fuel/fuel1.png" />
					<a href="/audi/2010/354/audi-a1-1.2-tfsi-sportback"
					   class="orange">
						Audi A1 Sportback 1.2 TFSI					</a>
				</div>
			</div>
			<div class="model-auto-podaci-data">
				<div class="model-auto-gorivo">
					Fuel:<span>Petrol</span>
				</div>
				<div class="model-auto-pogon">
					Wheel drive:<span>Front</span>
				</div>
				<div class="model-auto-kubikaza">
											Displacement:<span>1197 cm3</span>
									</div>
				<div class="model-auto-snaga">
					Power:<span>86 ks</span>
				</div>
			</div>
		</div>
			<div class="model-auto-podaci-wrap">
			<div class="model-auto-podaci-box">
				<div class="model-auto-puna-oznaka">
					<img class="fuel-badge" src="/resources/images/fuel/fuel1.png" />
					<a href="/audi/2010/355/audi-a1-1.4-tfsi-sportback"
					   class="orange">
						Audi A1 Sportback 1.4 TFSI					</a>
				</div>
			</div>
			<div class="model-auto-podaci-data">
				<div class="model-auto-gorivo">
					Fuel:<span>Petrol</span>
				</div>
				<div class="model-auto-pogon">
					Wheel drive:<span>Front</span>
				</div>
				<div class="model-auto-kubikaza">
											Displacement:<span>1390 cm3</span>
									</div>
				<div class="model-auto-snaga">
					Power:<span>122 ks</span>
				</div>
			</div>
		</div>
			<div class="model-auto-podaci-wrap">
			<div class="model-auto-podaci-box">
				<div class="model-auto-puna-oznaka">
					<img class="fuel-badge" src="/resources/images/fuel/fuel1.png" />
					<a href="/audi/2010/356/audi-a1-1.4-tfsi-cod-sportback"
					   class="orange">
						Audi A1 Sportback 1.4 TFSI COD					</a>
				</div>
			</div>
			<div class="model-auto-podaci-data">
				<div class="model-auto-gorivo">
					Fuel:<span>Petrol</span>
				</div>
				<div class="model-auto-pogon">
					Wheel drive:<span>Front</span>
				</div>
				<div class="model-auto-kubikaza">
											Displacement:<span>1390 cm3</span>
									</div>
				<div class="model-auto-snaga">
					Power:<span>140 ks</span>
				</div>
			</div>
		</div>
			<div class="model-auto-podaci-wrap">
			<div class="model-auto-podaci-box">
				<div class="model-auto-puna-oznaka">
					<img class="fuel-badge" src="/resources/images/fuel/fuel1.png" />
					<a href="/audi/2010/357/audi-a1-1.4-tfsi-185-sportback"
					   class="orange">
						Audi A1 Sportback 1.4 TFSI 185					</a>
				</div>
			</div>
			<div class="model-auto-podaci-data">
				<div class="model-auto-gorivo">
					Fuel:<span>Petrol</span>
				</div>
				<div class="model-auto-pogon">
					Wheel drive:<span>Front</span>
				</div>
				<div class="model-auto-kubikaza">
											Displacement:<span>1390 cm3</span>
									</div>
				<div class="model-auto-snaga">
					Power:<span>185 ks</span>
				</div>
			</div>
		</div>
			<div class="model-auto-podaci-wrap">
			<div class="model-auto-podaci-box">
				<div class="model-auto-puna-oznaka">
					<img class="fuel-badge" src="/resources/images/fuel/fuel2.png" />
					<a href="/audi/2010/358/audi-a1-1.6-tdi-sportback"
					   class="orange">
						Audi A1 Sportback 1.6 TDI					</a>
				</div>
			</div>
			<div class="model-auto-podaci-data">
				<div class="model-auto-gorivo">
					Fuel:<span>Diesel</span>
				</div>
				<div class="model-auto-pogon">
					Wheel drive:<span>Front</span>
				</div>
				<div class="model-auto-kubikaza">
											Displacement:<span>1598 cm3</span>
									</div>
				<div class="model-auto-snaga">
					Power:<span>90 ks</span>
				</div>
			</div>
		</div>
			<div class="model-auto-podaci-wrap">
			<div class="model-auto-podaci-box">
				<div class="model-auto-puna-oznaka">
					<img class="fuel-badge" src="/resources/images/fuel/fuel2.png" />
					<a href="/audi/2010/359/audi-a1-1.6-tdi-sportback"
					   class="orange">
						Audi A1 Sportback 1.6 TDI					</a>
				</div>
			</div>
			<div class="model-auto-podaci-data">
				<div class="model-auto-gorivo">
					Fuel:<span>Diesel</span>
				</div>
				<div class="model-auto-pogon">
					Wheel drive:<span>Front</span>
				</div>
				<div class="model-auto-kubikaza">
											Displacement:<span>1598 cm3</span>
									</div>
				<div class="model-auto-snaga">
					Power:<span>105 ks</span>
				</div>
			</div>
		</div>
			<div class="model-auto-podaci-wrap">
			<div class="model-auto-podaci-box">
				<div class="model-auto-puna-oznaka">
					<img class="fuel-badge" src="/resources/images/fuel/fuel2.png" />
					<a href="/audi/2010/360/audi-a1-2.0-tdi-sportback"
					   class="orange">
						Audi A1 Sportback 2.0 TDI					</a>
				</div>
			</div>
			<div class="model-auto-podaci-data">
				<div class="model-auto-gorivo">
					Fuel:<span>Diesel</span>
				</div>
				<div class="model-auto-pogon">
					Wheel drive:<span>Front</span>
				</div>
				<div class="model-auto-kubikaza">
											Displacement:<span>1968 cm3</span>
									</div>
				<div class="model-auto-snaga">
					Power:<span>143 ks</span>
				</div>
			</div>
		</div>
	</div>			</div>
		</div>
			</div>
	<div class="google-add-small">
		
	</div>
<div>
	<h2>
		<i class="fa fa-shield" aria-hidden="true"></i>
		Safety results for Audi A1	</h2>
</div>
<div class="box-wrap">
	<div class="euro-ncap-youtube">
		<div class="ncap-img">
			<img src="/resources/images/gui/details/euroncap.jpg">
		</div>
		<div class="stars">
			<div class="stars5"></div>
		</div>
		<div class="euro-ncap strong">
			EuroNCAP:
			90% occupant safety			
		</div>
					<iframe width="100%"
					height="360"
					src="https://www.youtube.com/embed/SOEuP95YxR4?rel=0"
					frameborder="0"
					allow="accelerometer; autoplay; encrypted-media; gyroscope; picture-in-picture"
					allowfullscreen>
			</iframe>
			</div>
</div><div>
	<h2><i class="fa fa-search" aria-hidden="true"></i>Check a car by its VIN number</h2>
</div>
<div class="box-wrap" style="height: 250px;">
	<div
	  data-cvaff
	  data-locale="en-EU"
	  data-a="automanijak"
	  data-b="81ec5429"
	  style="width: 100%; height: 100%">
	</div>
</div><div>
	<h2><i class="fa fa-user" aria-hidden="true"></i>Visitor comments</h2>
</div>
<div class="box-wrap">
	<div class="fb-comments" data-href="https://www.automaniac.orghttps://www.automaniac.org/model/audi/212/audi-a1-2010" data-width="670" data-numposts="5"></div>
</div>
	</div>
			</div>
		</div>
		<div id="footer">
			<div class="footer-wrap">
				Copyright &copy; 2015 - 2022 automaniac.org - All rights reserved.
				</br>
				Powered by <strong>Gama1 Solutions</strong>
				<div class="footer-links">
					<div class="footer-menu">
						<ul>
							<li><a href="/">HOME page</a></li>
							<li><a href="/about">about</a></li>
							<li><a href="/terms">terms & conditions</a></li>
							<li><a href="/legal">legal note</a></li>
							<li><a href="/faq">frequent questions (FAQ)</a></li>
							<li><a href="/contact">contact us</a></li>
						</ul>
					</div>

					<div class='footer-auto-brend'><ul><li><a href='/make/alfa-romeo'>Alfa Romeo</a></li><li><a href='/make/audi'>Audi</a></li><li><a href='/make/bmw'>BMW</a></li><li><a href='/make/chery'>Chery</a></li><li><a href='/make/chevrolet'>Chevrolet</a></li><li><a href='/make/chrysler'>Chrysler</a></li></ul></div><div class='footer-auto-brend'><ul><li><a href='/make/citroen'>Citroen</a></li><li><a href='/make/dacia'>Dacia</a></li><li><a href='/make/daihatsu'>Daihatsu</a></li><li><a href='/make/dodge'>Dodge</a></li><li><a href='/make/fiat'>FIAT</a></li><li><a href='/make/ford'>Ford</a></li></ul></div><div class='footer-auto-brend'><ul><li><a href='/make/honda'>Honda</a></li><li><a href='/make/hyundai'>Hyundai</a></li><li><a href='/make/infiniti'>Infiniti</a></li><li><a href='/make/jaguar'>Jaguar</a></li><li><a href='/make/jeep'>Jeep</a></li><li><a href='/make/kia'>KIA</a></li></ul></div><div class='footer-auto-brend'><ul><li><a href='/make/lada'>Lada</a></li><li><a href='/make/lancia'>Lancia</a></li><li><a href='/make/land-rover'>Land Rover</a></li><li><a href='/make/lexus'>Lexus</a></li><li><a href='/make/mazda'>Mazda</a></li><li><a href='/make/mercedes-benz'>Mercedes</a></li></ul></div><div class='footer-auto-brend'><ul><li><a href='/make/mini'>Mini</a></li><li><a href='/make/mitsubishi'>Mitsubishi</a></li><li><a href='/make/nissan'>Nissan</a></li><li><a href='/make/opel'>Opel</a></li><li><a href='/make/peugeot'>Peugeot</a></li><li><a href='/make/porsche'>Porsche</a></li></ul></div><div class='footer-auto-brend'><ul><li><a href='/make/renault'>Renault</a></li><li><a href='/make/rover'>Rover</a></li><li><a href='/make/saab'>SAAB</a></li><li><a href='/make/ssangyong'>SSangYong</a></li><li><a href='/make/seat'>Seat</a></li><li><a href='/make/smart'>Smart</a></li></ul></div><div class='footer-auto-brend'><ul><li><a href='/make/subaru'>Subaru</a></li><li><a href='/make/suzuki'>Suzuki</a></li><li><a href='/make/tesla'>Tesla</a></li><li><a href='/make/toyota'>Toyota</a></li><li><a href='/make/volkswagen'>Volkswagen</a></li><li><a href='/make/volvo'>Volvo</a></li></ul></div><div class='footer-auto-brend'><ul><li><a href='/make/zastava'>Zastava</a></li><li><a href='/make/škoda'>Škoda</a></li></ul></div>
				</div>
			</div>
		</div>
	</body>
</html>`

const version_html = `

<!DOCTYPE html>
<html lang="en">
	<head>
		<meta charset="utf-8">
		<title>
			Mazda CX-5  SkyActiv-D 2.2 175 4WD (2017 - ) - AutoManiac		</title>
		<meta content="width=device-width, initial-scale=1.0" name="viewport">
		<meta content="2017 Mazda CX-5 ,specifications,information,dimensions,consumption,acceleration,top speed,fuel tank volume,reliability,failures,opinions,safety,crash test" name="keywords">
		<meta content="2017 Mazda CX-5  SkyActiv-D 2.2 175 4WD - Information, specifications and drivers opinions" name="description">
		<meta http-equiv="Cache-Control" content="no-cache, max-age=0, must-revalidate, no-store">
		<meta property="og:site_name" content="AutoManiac"/>
		<meta property="og:image" content="/resources/images/variant/1940/cx-5_2.jpg" />
		<meta property="og:image:width" content="700"/>
		<meta property="og:image:height" content="300"/>
		<meta property="og:description" content="2017 Mazda CX-5  SkyActiv-D 2.2 175 4WD - Information, specifications and drivers opinions" />
		<meta property="og:title" content="Mazda CX-5  SkyActiv-D 2.2 175 4WD (2017 - )" />
		<meta property="fb:app_id" content="290604621342699" />
		<meta property="fb:admins" content="1276696339"/>
		<link rel="image_src" type="image/jpeg" href="/resources/images/variant/1940/cx-5_2.jpg" />
		<link rel="shortcut icon" href="/resources/images/gui/favicon.png" type="image/x-icon">
		<link rel="icon" href="/resources/images/gui/favicon.png" type="image/x-icon">
			<script type="text/javascript" src="/resources/js/jquery-1.8.3.js"></script>
	<script type="text/javascript" src="/resources/js/jquery-ui.js"></script>
	<script type="text/javascript" src="/resources/js/scripts.js?version=v2.1.0"></script>
	<script type="text/javascript" src="/resources/js/jquery.chosen.min.js"></script>
	<script src="https://apis.google.com/js/platform.js" async defer></script>
	<script async src="https://securepubads.g.doubleclick.net/tag/js/gpt.js"></script>
	<!--
	<script async src="//pagead2.googlesyndication.com/pagead/js/adsbygoogle.js"></script>
	<script type="text/javascript" src="/resources/js/adsense.js"></script>
	-->
	<script async defer crossorigin="anonymous" src="https://connect.facebook.net/en_GB/sdk.js#xfbml=1&version=v4.0"></script>
	
						
	<link rel="stylesheet" href="/resources/css/styles.css?version=v2.1.6" type="text/css" media="screen" charset="utf-8" />
	<link rel="stylesheet" href="/resources/css/chosen.css" type="text/css" media="screen" charset="utf-8" />
	<link rel="stylesheet" href="/resources/css/font-awesome.min.css" >
	<link rel="stylesheet" href="https://fonts.googleapis.com/css?family=Open+Sans:300,300i,400,400i,700,700i|Raleway:300,400,500,700,800|Montserrat:300,400,700">
		<link rel="icon" href="/resources/images/gui/favicon.png" >
		<link rel="shortcut icon" href="/resources/images/gui/favicon.png" type="image/x-icon">
		<link rel="apple-touch-icon" href="images/gui/apple-touch-icon.png">
	</head>

	<body root-url="">

		<!-- Load Facebook SDK for JavaScript -->
		<div id="fb-root"></div>
		<script>
		  window.fbAsyncInit = function() {
		    FB.init({
		      appId      : '290604621342699',
		      xfbml      : true,
		      version    : 'v2.8'
		    });
		  };

		  (function(d, s, id){
		     var js, fjs = d.getElementsByTagName(s)[0];
		     if (d.getElementById(id)) {return;}
		     js = d.createElement(s); js.id = id;
		     js.src = "//connect.facebook.net/en_US/sdk.js";
		     fjs.parentNode.insertBefore(js, fjs);
		   }(document, 'script', 'facebook-jssdk'));

		  (function(w,d,u,h,s){
		    h=d.getElementsByTagName('head')[0];
		    s=d.createElement('script');
		    s.async=1;
		    s.src=u+'/sdk.js';
		    h.appendChild(s);
		  })(window,document,'https://aff.carvertical.com');

			
			  (function(i,s,o,g,r,a,m){i['GoogleAnalyticsObject']=r;i[r]=i[r]||function(){
			  (i[r].q=i[r].q||[]).push(arguments)},i[r].l=1*new Date();a=s.createElement(o),
			  m=s.getElementsByTagName(o)[0];a.async=1;a.src=g;m.parentNode.insertBefore(a,m)
			  })(window,document,'script','//www.google-analytics.com/analytics.js','ga');
			
			  ga('create', 'UA-75322623-2', 'auto');
	  		  ga('send', 'pageview');

	  		
		
		</script>

		<div class="menu-wrap">
			<div class="topnav" id="myTopnav">
				<div class="top-bar">
					<div class="top-bar-wrap">
						<div class="top-bar-left">
							<a href="/about" title="about">about</a>
							<a href="/legal" title="legal note">legal note</a>
							<a href="/contact" title="contact us">contact us</a>
							<span style="color: #fff; font-weight: bold;"></span>
						</div>
						<div class="top-bar-right">
							<a href="#" title="">
								Login <i class="fa fa-user" aria-hidden="true"></i>
							</a>
							<a href="https://www.automanie.org" title="">
								deutsch (D, AT, CH)<i class="fa fa-globe" aria-hidden="true"></i>
							</a>
							<a href="https://www.automanijak.com" title="">
								srpski (ex-yugoslavia)<i class="fa fa-flag" aria-hidden="true"></i>
							</a>
						</div>
					</div>
				</div>
				<a href="#" class="icon"><i class="fa fa-bars"></i></a>
				<div class="bottom-bar">
					<div class="bottom-bar-wrap">
						<div class="bottom-bar-left">
							<div class="logo">
								<a href="/" title="">
									<img src="/resources/images/gui/logo.png" alt="" title="">
								</a>
							</div>
						</div>
						<div class="bottom-bar-right">
							<a href="/virtual-adviser">Virtual adviser</a>
							<a href="/specs">Specifications by model</a>
							<a href="/compare">Compare two cars</a>
							<a href="/badges">Automotive badges</a>
							<a href="/blog">Blog</a>
						</div>
					</div>
				</div>
			</div>
		</div>
		<div id="main-wrapper">
			<div class="left-ad">
				<script async src="//pagead2.googlesyndication.com/pagead/js/adsbygoogle.js"></script>
<!-- branding -->
<ins class="adsbygoogle branding-ad"
     style="width:300px;height:600px"
     data-ad-client="ca-pub-2798966384669812"
     data-ad-slot="1676022217"></ins>
<script>
(adsbygoogle = window.adsbygoogle || []).push({});
</script>				<script async src="https://pagead2.googlesyndication.com/pagead/js/adsbygoogle.js"></script>
<!-- massive-header -->
<ins class="adsbygoogle massive-header-ad"
     data-ad-client="ca-pub-2798966384669812"
     data-ad-slot="9878901748"
     data-full-width-responsive="true"></ins>
<script>
     (adsbygoogle = window.adsbygoogle || []).push({});
</script>			</div>
			<div class="right-ad">
				<script async src="//pagead2.googlesyndication.com/pagead/js/adsbygoogle.js"></script>
<!-- branding -->
<ins class="adsbygoogle branding-ad"
     style="width:300px;height:600px"
     data-ad-client="ca-pub-2798966384669812"
     data-ad-slot="1676022217"></ins>
<script>
(adsbygoogle = window.adsbygoogle || []).push({});
</script>			</div>
			<div id="wrapper" class="clearfix">	<div id="sidebar">
		<div class="sidebar-wrap">
		<div id="sidebar-dodatno" class="search-box">
	<div class="sidebar-dodatno-wrap">
		<script async src="https://cse.google.com/cse.js?cx=016645478549947938819:d5ilzoxfr9e"></script>
		<script>
			window.onload = function(){
				$(".gsc-input").attr("placeholder", "Search this website...");
			}
		</script>
		<div class="gcse-search"></div>
	</div>
</div>	
	<div id="sidebar-dodatno">
		<div class="sidebar-misljenje-naslov">
			<h2>Have you driven it?</h2>
		</div>
		<div class="sidebar-dodatno-wrap">
			<div class="misljenje-vozaca">
				<span>Reliability <strong title="- Malfunction frequency&#13;- Malfunction severity&#13;- Spare parts availability">(?)</strong>:</span>
				<div class="avg" related-field="1">
					<img src='/resources/images/gui/details/star-black-empty.png'
						 class='rating-star'
						 car-id='12927'
						 related-field='1'
						 star-no='1'/>
					 <img src='/resources/images/gui/details/star-black-empty.png'
						 class='rating-star'
						 car-id='12927'
						 related-field='1'
						 star-no='2'/>
					<img src='/resources/images/gui/details/star-black-empty.png'
						 class='rating-star'
						 car-id='12927'
						 related-field='1'
						 star-no='3'/>
					<img src='/resources/images/gui/details/star-black-empty.png'
						 class='rating-star'
						 car-id='12927'
						 related-field='1'
						 star-no='4'/>
					<img src='/resources/images/gui/details/star-black-empty.png'
						 class='rating-star'
						 car-id='12927'
						 related-field='1'
						 star-no='5'/>
					<div class="misljenje-ocena">
					?					</div>
				</div>
				<div id="issues" class="hidden">
					<form>
						<br/>
						Any engine issues?						<br/><br/>
						<div class='selectable' issue_id='1' engine_id='794'>
							<div class='class-description'>mechanical / belts</div>
						</div>
						<div class='selectable' issue_id='2' engine_id='794'>
							<div class='class-description'>electrical / ECU</div>
						</div>
						<div class='selectable' issue_id='3' engine_id='794'>
							<div class='class-description'>fuel supply</div>
						</div>
						<div class='selectable' issue_id='4' engine_id='794'>
							<div class='class-description'>air supply</div>
						</div>
						<div class='selectable' issue_id='5' engine_id='794'>
							<div class='class-description'>cooling system</div>
						</div>
						<div class='selectable' issue_id='6' engine_id='794'>
							<div class='class-description'>clutch / flywheel</div>
						</div>
						<div class='selectable' issue_id='7' engine_id='794'>
							<div class='class-description'>exchaust (EGR / DPF)</div>
						</div>
						<br/>
						<div id="report-issue" class='submit-issues' engine_id='794'>
							<i class="fa fa-wrench" aria-hidden="true"></i>
							<span>Report</span>
						</div>
					</form>
				</div>
			</div>
			<div class="misljenje-vozaca">
				<span>Performance <strong title="- Acceleration&#13;- Top speed&#13;- Braking efficiency&#13;- Handling">(?)</strong>:</span>
				<div class="avg" related-field="2">
					<img src='/resources/images/gui/details/star-black-empty.png'
						 class='rating-star'
						 car-id='12927'
						 related-field='2'
						 star-no='1'/>
					 <img src='/resources/images/gui/details/star-black-empty.png'
						 class='rating-star'
						 car-id='12927'
						 related-field='2'
						 star-no='2'/>
					<img src='/resources/images/gui/details/star-black-empty.png'
						 class='rating-star'
						 car-id='12927'
						 related-field='2'
						 star-no='3'/>
					<img src='/resources/images/gui/details/star-black-empty.png'
						 class='rating-star'
						 car-id='12927'
						 related-field='2'
						 star-no='4'/>
					<img src='/resources/images/gui/details/star-black-empty.png'
						 class='rating-star'
						 car-id='12927'
						 related-field='2'
						 star-no='5'/>
					<div class="misljenje-ocena">
					?					</div>
				</div>
			</div>
			<div class="misljenje-vozaca">
				<span>Economy <strong title="- Fuel consumption&#13;- Regular maintenance prices&#13;- Spare parts prices&#13;- Other expenses specific to this car">(?)</strong>:</span>
				<div class="avg" related-field="3">
					<img src='/resources/images/gui/details/star-black-empty.png'
						 class='rating-star'
						 car-id='12927'
						 related-field='3'
						 star-no='1'/>
					 <img src='/resources/images/gui/details/star-black-empty.png'
						 class='rating-star'
						 car-id='12927'
						 related-field='3'
						 star-no='2'/>
					<img src='/resources/images/gui/details/star-black-empty.png'
						 class='rating-star'
						 car-id='12927'
						 related-field='3'
						 star-no='3'/>
					<img src='/resources/images/gui/details/star-black-empty.png'
						 class='rating-star'
						 car-id='12927'
						 related-field='3'
						 star-no='4'/>
					<img src='/resources/images/gui/details/star-black-empty.png'
						 class='rating-star'
						 car-id='12927'
						 related-field='3'
						 star-no='5'/>
					<div class="misljenje-ocena">
					?					</div>
				</div>
			</div>
			<div class="misljenje-vozaca">
				<span>Comfort <strong title="- Suspension quality&#13;- Seats comfort&#13;- Legs and head space&#13;- Driver's position & ergonomy">(?)</strong>:</span>
				<div class="avg" related-field="4">
					<img src='/resources/images/gui/details/star-black-empty.png'
						 class='rating-star'
						 car-id='12927'
						 related-field='4'
						 star-no='1'/>
					 <img src='/resources/images/gui/details/star-black-empty.png'
						 class='rating-star'
						 car-id='12927'
						 related-field='4'
						 star-no='2'/>
					<img src='/resources/images/gui/details/star-black-empty.png'
						 class='rating-star'
						 car-id='12927'
						 related-field='4'
						 star-no='3'/>
					<img src='/resources/images/gui/details/star-black-empty.png'
						 class='rating-star'
						 car-id='12927'
						 related-field='4'
						 star-no='4'/>
					<img src='/resources/images/gui/details/star-black-empty.png'
						 class='rating-star'
						 car-id='12927'
						 related-field='4'
						 star-no='5'/>
					<div class="misljenje-ocena">
					?					</div>
				</div>
			</div>
			<div class="misljenje-vozaca">
				<span>Practicality <strong title="- Easy to use&#13;- Quantity & quality of in-car compartments&#13;- Luggage space access&#13;- Spare parts replacement ease">(?)</strong>:</span>
				<div class="avg" related-field="5">
					<img src='/resources/images/gui/details/star-black-empty.png'
						 class='rating-star'
						 car-id='12927'
						 related-field='5'
						 star-no='1'/>
					 <img src='/resources/images/gui/details/star-black-empty.png'
						 class='rating-star'
						 car-id='12927'
						 related-field='5'
						 star-no='2'/>
					<img src='/resources/images/gui/details/star-black-empty.png'
						 class='rating-star'
						 car-id='12927'
						 related-field='5'
						 star-no='3'/>
					<img src='/resources/images/gui/details/star-black-empty.png'
						 class='rating-star'
						 car-id='12927'
						 related-field='5'
						 star-no='4'/>
					<img src='/resources/images/gui/details/star-black-empty.png'
						 class='rating-star'
						 car-id='12927'
						 related-field='5'
						 star-no='5'/>
					<div class="misljenje-ocena">
					?					</div>
				</div>
			</div>
		</div>
	</div><div id="sidebar-dodatno">
	<div class="sidebar-sve-varijante-naslov">
		<h2>All model's variants</h2>
	</div>
	<div class="sidebar-dodatno-wrap">
			
	<div class="najpop-item">
		<div class="najpop-item-name">
			<a href="/model/mazda/1285/mazda-cx-5-2017" title="2017 Mazda CX-5 ">
				2017 Mazda CX-5 			</a>
		</div>
		<div class="najpop-item-img">
			<a href="/model/mazda/1285/mazda-cx-5-2017" title="2017 Mazda CX-5 ">
				<img src="/resources/images/variant/1940/thumb_cx-5_2.jpg" 	title="Mazda CX-5"
					 alt="Mazda CX-5"
				 	 class="thumb-image">
			</a>
			<div class="ocena"><span>4.8</span>
			</div>
		</div>
		<div class="najpop-item-val">
			from: 			</br>
			<strong>23.290</strong>
			</br>EUR</div>
	</div>	</div>
</div><div id="sidebar-dodatno">
	<div class="sidebar-minimalna-cena-naslov">
		<h2>Starting price</h2>
	</div>
	<div class="sidebar-dodatno-wrap">
		<div class="dodatno-text">
			<div>
				from				</br>
				<strong>23.290</strong>
				</br>
				EUR*
			</div>
			<span>Starting price is informative only and is being updated using information from leading online yellow pages. Prices of well maintained examples are commonly 20% higher.</span> 
		</div>
	</div>
</div>		</div>
		<div class="side-ad">
	<script async src="//pagead2.googlesyndication.com/pagead/js/adsbygoogle.js"></script>
	<ins class="adsbygoogle side-ad-vertical"
	     style="width:160px;height:600px"
	     data-ad-client="ca-pub-2798966384669812"
	     data-ad-slot="9565981886"></ins>
	<script>
	     (adsbygoogle = window.adsbygoogle || []).push({});
	</script>

	<ins class="adsbygoogle side-ad-horizontal"
	     data-ad-client="ca-pub-2798966384669812"
	     data-ad-slot="8347642282"
	     data-full-width-responsive="true"></ins>
	<script>
	     (adsbygoogle = window.adsbygoogle || []).push({});
	</script>
</div>
	</div>
	<div id="page-wrap">
<div id="breadcrumb-wrap">
	<div class="breadcrumb-nav"><a href='/'>HOME page</a> / <a href='/specs'>specifications</a> / <a href='/make/mazda'>Mazda</a> / <a href='/model/mazda/1285/mazda-cx-5-2017'>2017 CX-5</a> /  SkyActiv-D 2.2 175 4WD</div>
	<div id="sidebar-dodatno" class="search-box">
	<div class="sidebar-dodatno-wrap">
		<script async src="https://cse.google.com/cse.js?cx=016645478549947938819:d5ilzoxfr9e"></script>
		<script>
			window.onload = function(){
				$(".gsc-input").attr("placeholder", "Search this website...");
			}
		</script>
		<div class="gcse-search"></div>
	</div>
</div>	<div class="facebook">
				<div class="fb-like" data-href="https://www.automaniac.org/mazda/2017/12927/mazda-cx-5-skyactiv-d-2.2-175-4wd" data-width="" data-layout="button_count" data-action="like" data-size="small" data-show-faces="true" data-share="true"></div>
	</div>
</div>
	<div id="predlog-auto">
		<div class="box-wrap">
			<div class="predlog-auto-logo">
				<a href="/make/mazda">
					<img src="/resources/images/make/mazda.jpg" title="Mazda" />
				</a>
			</div>
			<div class="predlog-auto-naslov">
				Mazda CX-5  SkyActiv-D 2.2 175 4WD				<div class="predlog-auto-stats">
					Model Year 2017, J - Segment (SUV)				</div>
			</div>
			<div class="predlog-auto-slika">
				<div class="auto-score-wrap">
										<div class="score-left">
						<div class="score3b">
							<p>drivers'</p>
							<span>3.0</span>
							<p>rating</p>
						</div>
					</div>
					<div class="score-right clearfix">
						<div class="score5a"></div>
						<div class="score4a"></div>
						<div class="score3a"></div>
						<div class="score2a"></div>
						<div class="score1a"></div>
					</div>
				</div>
				<img src="/resources/images/variant/1940/cx-5_2.jpg" />
										<div class="predlog-auto-oznaka-wrap">
							<span>Badges:</span>
															<div class="predlog-auto-oznaka">
									<a href="/badges/85/mazda-skyactiv">
										<img src="/resources/images/badge/mazda_skyactiv.jpg"
											 title="Mazda - SKYACTIV badge"
											 alt="Mazda - SKYACTIV" />
									</a>
								</div>
													</div>
								<div class="predlog-auto-uporedi clearfix">
					<a href='/compare/12927/0/mazda-cx-5-skyactiv-d-2.2-175-4wd-VS-'>
						<span>Compare with</span>
						<i class="fa fa-exchange" aria-hidden="true"></i>
					</a>
				</div>
			</div>
		</div>
		<div class="box-wrap">
			<div class="predlog-auto-box-wrap">
				<div class="predlog-auto-detalj" title="Body style">
					<div class="predlog-auto-detalj-slika">
						<img src="/resources/images/gui/details/body_6.jpg"
							 alt="suv" />
					</div>
					<span>suv</span>
				</div>
				<div class="predlog-auto-detalj" title="Doors No.">
					<div class="predlog-auto-detalj-slika">
						<img src="/resources/images/gui/details/door.jpg"
							 alt="x 5" />
					</div>
					<span>x 5</span>
				</div>
				<div class="predlog-auto-detalj" title="Wheel drive">
										<div class="predlog-auto-detalj-slika">
						<img src="/resources/images/gui/details/drive_3.jpg"
							 alt="4 x 4" />
					</div>
					<span>
					4 x 4					</span>
				</div>
				<div class="predlog-auto-detalj" title="Engine configuration">
					<div class="predlog-auto-detalj-slika">
						<img src="/resources/images/gui/details/configuration.jpg" alt="" />
					</div>
					<span>
						2.2 L4 16v					</span>
				</div>
									<div class="predlog-auto-detalj" title="Aspiration">
						<div class="predlog-auto-detalj-slika">
							<img src="/resources/images/gui/details/turbo.jpg"
								 alt="Turbo" />
						</div>
						<span>Turbo</span>
					</div>
								<div class="predlog-auto-detalj" title="Fuel type">
					<div class="predlog-auto-detalj-slika">
						<img src="/resources/images/gui/details/fuel.jpg"
							 alt="Diesel">
					</div>
					<span>Diesel</span>
				</div>
			</div>
		</div>
<div>
	<h2><i class="fa fa-info-circle" aria-hidden="true"></i>Basic info on Mazda CX-5  SkyActiv-D 2.2 175 4WD</h2>
</div>
<div class="box-wrap">
	<p>
		The Japanese car was first shown in year 2017 and powered by a 4 - cylinder turbo diesel unit, produced by Mazda. The engine offers a displacement of  2.2 litre matched to a 4 x 4 wheel drive system and a manual gearbox with 6  or a automatic gearbox with 6 gears. Vehicle in question is a suv with the top speed of 207km/h, reaching the 100km/h (62mph) mark in 8.8s and consuming around 5.2 liters of fuel every 100 kilometers.
	</p>
</div>		<div>
			<h2><i class="fa fa-database" aria-hidden="true"></i>Data</h2>
		</div>
		<div class="podaci-box-wrap">
<div class="podaci-box-left-wrap">
	<div class="podaci-naslov">
	Engine info	(<a href="/engine/mazda/794/mazda-2191cc-diesel-2.2-skyactiv-d-16v-175hp" class="orange">
		 issues & breakdowns 	</a>)
		</div>
	<div class="podaci-box-a">
		<a href="/make/mazda">
			<img src="/resources/images/make/mazda.jpg"
				 title="Mazda" />
		</a>
	</div>
	<div class="podaci-box-b">
		Mazda		<p>
							<a href="/engine/mazda/794/mazda-2191cc-diesel-2.2-skyactiv-d-16v-175hp" class="orange">2.2 SKYACTIV-D</a>
					</p>
	</div>
	<div class="podaci-box-c">
			</div>
			<div class="data">
			<div class="d1">displacement:</div>
			<div class="d2"><strong>2191</strong></div>
			<div class="d3">cc</div>
		</div>
		<div class="data">
			<div class="d1">configuration:</div>
			<div class="d2"><strong>4  - Inline</strong></div>
			<div class="d3">&nbsp;</div>
		</div>
		<div class="data">
			<div class="d1">valves:</div>
			<div class="d2"><strong>16, 4 p/ cylinder</strong></div>
			<div class="d3">&nbsp;</div>
		</div>
		<div class="data">
			<div class="d1">aspiration:</div>
			<div class="d2"><strong>Turbo</strong></div>
		<div class="d3">&nbsp;</div>
	</div>
		<div class="data">
		<div class="d1">fuel type:</div>
		<div class="d2"><strong>Diesel</strong></div>
		<div class="d3">&nbsp;</div>
	</div>
	<div class="data">
		<div class="d1">power:</div>
		<div class="d2"><strong>175</strong></div>
		<div class="d3">hp</div>
	</div>
	<div class="data">
		<div class="d1">torque:</div>
		<div class="d2"><strong>420</strong>
		</div>
		<div class="d3">Nm</div>
	</div>
</div><div class="podaci-box-right-wrap">
	<div class="podaci-naslov">Vehicle information - Mazda CX-5</div>
	<div class="podaci-box-a pola">
		<div class="podaci-euro-ncap">
			<img src="/resources/images/gui/details/euro-ncap.png">
		</div>
	</div>
	<div class="podaci-box-b pola">
		<div class="preporuceni-stars">
			<div class="stars5"></div>
			<div class="preporuceni-euro-ncap strong">
				94% score			</div>
		</div>
	</div>
	<div class="data">
		<div class="d1">produced:</div>
		<div class="d2"><strong>2017. - </strong></div>
		<div class="d3">&nbsp;</div>
	</div>
	<div class="data">
		<div class="d1">body style:</div>
		<div class="d2"><strong>suv</strong></div>
		<div class="d3">5 door</div>
	</div>
	<div class="data">
		<div class="d1">length:</div>
		<div class="d2"><strong>4555</strong></div>
		<div class="d3">mm</div>
	</div>
	<div class="data">
		<div class="d1">width:</div>
		<div class="d2"><strong>1840</strong></div>
		<div class="d3">mm</div>
	</div>
	<div class="data">
		<div class="d1">height:</div>
		<div class="d2"><strong>1710</strong>
		</div>
		<div class="d3">mm</div>
	</div>
	<div class="data">
		<div class="d1">boot:</div>
		<div class="d2"><strong>503 - 1620</strong>
		</div>
		<div class="d3">l</div>
	</div>
			<div class="data">
			<div class="d1">fuel tank:</div>
			<div class="d2"><strong>58</strong>
			</div>
			<div class="d3">l</div>
		</div>
	</div>		</div>
		<div>
			<h2><i class="fa fa-signal" aria-hidden="true"></i>Performance</h2>
		</div>
		<div class="podaci-box-wrap">
<div class="podaci-box-left-wrap">
			<div class="podaci-naslov">Manual gearbox performance</div>
		<div class="podaci-box-a">
			<div class="energy-left ">
				<div class="energyAb"></div>
				<div class="energyBb"></div>
				<div class="energyCb"></div>
				<div class="energyDb"></div>
				<div class="energyEb"></div>
				<div class="energyFb"></div>
				<div class="energyGb"></div>
			</div>
			<div class="energy-right">
				<div class="energyCa">
					<p>emission</p>
					<span>C</span>
					<p>category</p>
				</div>
			</div>
		</div>
		<div class="podaci-box-b">
			manual gearbox			<p>6 gears</p>
		</div>
		<div class="podaci-box-c">
			<div class="podaci-prosek">
				<p>
				average				</p>
				<span>
				5.2				</span>
				<p>
				l/100km				</p>
			</div>
		</div>
		<div class="data">
			<div class="d1">weight:</div>
			<div class="d2"><strong>1480</strong></div>
			<div class="d3">kg</div>
		</div>
		<div class="data">
			<div class="d1">acc. 0-100 km/h:</div>
			<div class="d2"><strong>8.8</strong></div>
			<div class="d3">s</div>
		</div>
		<div class="data">
			<div class="d1">top speed:</div>
			<div class="d2"><strong>207</strong></div>
			<div class="d3">km/h</div>
		</div>
					<div class="data">
				<div class="d1">cons. (urban):</div>
				<div class="d2"><strong>6.0</strong></div>
				<div class="d3">l/100km</div>
			</div>
			<div class="data">
				<div class="d1">cons. (highway):</div>
				<div class="d2"><strong>4.7</strong></div>
				<div class="d3">l/100km</div>
			</div>
			<div class="data">
				<div class="d1">cons. (average):</div>
				<div class="d2"><strong>5.2</strong></div>
				<div class="d3">l/100km</div>
			</div>
			<div class="data">
				<div class="d1">CO2 emissions:</div>
				<div class="d2"><strong>136</strong></div>
				<div class="d3">g/km</div>
			</div>
		</div>
<div class="podaci-box-right-wrap">
			<div class="podaci-naslov">Automatic gearbox performance</div>
		<div class="podaci-box-a">
			<div class="energy-left ">
				<div class="energyAb"></div>
				<div class="energyBb"></div>
				<div class="energyCb"></div>
				<div class="energyDb"></div>
				<div class="energyEb"></div>
				<div class="energyFb"></div>
				<div class="energyGb"></div>
			</div>
			<div class="energy-right">
				<div class="energyDa">
					<p>emission</p>
					<span>D</span>
					<p>category</p>
				</div>
			</div>
		</div>
		<div class="podaci-box-b">
			automatic			<p>6 gears</p>
		</div>
		<div class="podaci-box-c">
			<div class="podaci-prosek">
				<p>
				average				</p>
				<span>
				5.5				</span>
				<p>
				l/100km				</p>
			</div>
		</div>
		<div class="data">
			<div class="d1">weight:</div>
			<div class="d2"><strong>1495</strong></div>
			<div class="d3">kg</div>
		</div>
		<div class="data">
			<div class="d1">acc. 0-100 km/h:</div>
			<div class="d2"><strong>9.4</strong></div>
			<div class="d3">s</div>
		</div>
		<div class="data">
			<div class="d1">top speed:</div>
			<div class="d2"><strong>204</strong></div>
			<div class="d3">km/h</div>
		</div>
					<div class="data">
				<div class="d1">cons. (urban):</div>
				<div class="d2"><strong>6.4</strong></div>
				<div class="d3">l/100km</div>
			</div>
			<div class="data">
				<div class="d1">cons. (highway):</div>
				<div class="d2"><strong>4.9</strong></div>
				<div class="d3">l/100km</div>
			</div>
			<div class="data">
				<div class="d1">cons. (average):</div>
				<div class="d2"><strong>5.5</strong></div>
				<div class="d3">l/100km</div>
			</div>
			<div class="data">
				<div class="d1">CO2 emissions:</div>
				<div class="d2"><strong>144</strong></div>
				<div class="d3">g/km</div>
			</div>
		</div>
		</div>
	</div>
<div>
	<h2><i class="fa fa-sort" aria-hidden="true"></i>Pros & Cons compared to direct rivals</h2>
</div>
<div class="box-wrap">
	
	<div class="prednosti-nedostaci">
		<div class="good">
			safety		</div>
	</div>

			<div class="prednosti-nedostaci">
			<div class="good">
				reliability			</div>
		</div>
	
			<div class="prednosti-nedostaci">
			<div class="good">
				comfort			</div>
		</div>
	
			<div class="prednosti-nedostaci">
			<div class="good">
				practicality			</div>
		</div>
	
			<div class="prednosti-nedostaci">
			<div class="good">
				performance			</div>
		</div>
	
			<div class="prednosti-nedostaci">
			<div class="good">
				economy			</div>
		</div>
	
			<div class="prednosti-nedostaci">
			<div class="bad">
				boot space			</div>
		</div>
	
	
	
	
	
	
	
	
</div><div>
	<h2>
		<i class="fa fa-shield" aria-hidden="true"></i>
		Safety results for Mazda CX-5	</h2>
</div>
<div class="box-wrap">
	<div class="euro-ncap-youtube">
		<div class="ncap-img">
			<img src="/resources/images/gui/details/euroncap.jpg">
		</div>
		<div class="stars">
			<div class="stars5"></div>
		</div>
		<div class="euro-ncap strong">
			EuroNCAP:
			94% occupant safety			
		</div>
					<iframe width="100%"
					height="360"
					src="https://www.youtube.com/embed/BeAqCuLeld0?rel=0"
					frameborder="0"
					allow="accelerometer; autoplay; encrypted-media; gyroscope; picture-in-picture"
					allowfullscreen>
			</iframe>
			</div>
</div>	<div>
		<h2><i class="fa fa-exchange" aria-hidden="true"></i>Other cars that might interest you...</h2>
	</div>
<div class="box-wrap">
	<div class="preporuka-auto">
		<div class="preporuka-left">
			<div class="preporuka-slika">
				<a href="/compare/12927/11777/mazda-cx-5-skyactiv-d-2.2-175-4wd-VS-kia-sportage-2.0-crdi-4wd" title="compare selected cars">
					<img src="/resources/images/variant/1778/thumb_sportage_5.jpg" class="thumb-image">
				</a>
				<div class="preporuka-uporedi">
					<a href="/compare/12927/11777/mazda-cx-5-skyactiv-d-2.2-175-4wd-VS-kia-sportage-2.0-crdi-4wd" title="compare selected cars"><i class="fa fa-exchange" aria-hidden="true"></i></a>
				</div>
			</div>
		</div>
		<div class="preporuka-right">
			<div class="preporuka-naslov">
				<a href="/kia/2016/11777/kia-sportage-2.0-crdi-4wd" class="orange">KIA Sportage  2.0 CRDi 4WD</a>
			</div>
			<div class="preporuka-proizvodnja">
				produced from 2016. to 2018. &nbsp;			</div>
			<div class="preporuceni-stars">
				<div class="stars5"></div>
			</div>
			<div class="preporuceni-euro-ncap">
				EuroNCAP:
				90% score			</div>
		</div>
		<div class="preporuka-data-wrap">
			<div class="preporuka-data">
				<p>acceleration</p>
				<p>
					<strong>
						<div class="data-bad">
						7%						</div>
					</strong>
				</p>
				<p>slower</p>
			</div>
			<div class="preporuka-data">
				<p>consumption</p>
				<p>
					<strong>
						<div class="data-bad">
						12%						</div>
					</strong>
				</p>
				<p>higher</p>
			</div>
			<div class="preporuka-data">
				<p>power</p>
				<p>
					<strong>
						<div class="data-good">
						5%						</div>
					</strong>
				</p>
				<p>higher</p>
			</div>
			<div class="preporuka-data">
				<p>length</p>
				<p>
					<strong>
						<div class="data-bad">
						2%						</div>
					</strong>
				</p>
				<p>shorter</p>
			</div>
			<div class="preporuka-data">
				<p>weight</p>
				<p>
					<strong>
						<div class="data-bad">
						16%						</div>
					</strong>
				</p>
				<p>higher</p>
			</div>
							<div class="preporuka-data">
					<p>fuel tank</p>
					<p>
						<strong>
							<div class="data-good">
							6%							</div>
						</strong>
					</p>
					<p>larger</p>
				</div>
						<div class="preporuka-data">
				<p>boot</p>
				<p>
					<strong>
						<div class="data-bad">
						2%						</div>
					</strong>
				</p>
				<p>smaller</p>
			</div>
						<div class="preporuka-data">
				<p>boot ext.</p>
				<p>
					<strong>
						<div class="data-bad">
						9%						</div>
					</strong>
				</p>
				<p>smaller</p>
			</div>		
						<div class="preporuka-data">
				<p>price</p>
				<p>
					<strong>
						<div class="data-good">
						46%						</div>
					</strong>
				</p>
				<p>lower</p>
			</div>
		</div>
	</div>
</div><div class="box-wrap">
	<div class="preporuka-auto">
		<div class="preporuka-left">
			<div class="preporuka-slika">
				<a href="/compare/12927/11964/mazda-cx-5-skyactiv-d-2.2-175-4wd-VS-toyota-rav4-2.0-d-4d" title="compare selected cars">
					<img src="/resources/images/variant/1792/thumb_rav4_8.jpg" class="thumb-image">
				</a>
				<div class="preporuka-uporedi">
					<a href="/compare/12927/11964/mazda-cx-5-skyactiv-d-2.2-175-4wd-VS-toyota-rav4-2.0-d-4d" title="compare selected cars"><i class="fa fa-exchange" aria-hidden="true"></i></a>
				</div>
			</div>
		</div>
		<div class="preporuka-right">
			<div class="preporuka-naslov">
				<a href="/toyota/2016/11964/toyota-rav4-2.0-d-4d" class="orange">Toyota RAV4  2.0 D-4D</a>
			</div>
			<div class="preporuka-proizvodnja">
				produced from 2016. to 2018. &nbsp;			</div>
			<div class="preporuceni-stars">
				<div class="stars5"></div>
			</div>
			<div class="preporuceni-euro-ncap">
				EuroNCAP:
				89% score			</div>
		</div>
		<div class="preporuka-data-wrap">
			<div class="preporuka-data">
				<p>acceleration</p>
				<p>
					<strong>
						<div class="data-bad">
						8%						</div>
					</strong>
				</p>
				<p>slower</p>
			</div>
			<div class="preporuka-data">
				<p>consumption</p>
				<p>
					<strong>
						<div class="data-good">
						11%						</div>
					</strong>
				</p>
				<p>lower</p>
			</div>
			<div class="preporuka-data">
				<p>power</p>
				<p>
					<strong>
						<div class="data-bad">
						22%						</div>
					</strong>
				</p>
				<p>lower</p>
			</div>
			<div class="preporuka-data">
				<p>length</p>
				<p>
					<strong>
						<div class="data-good">
						1%						</div>
					</strong>
				</p>
				<p>longer</p>
			</div>
			<div class="preporuka-data">
				<p>weight</p>
				<p>
					<strong>
						<div class="data-bad">
						8%						</div>
					</strong>
				</p>
				<p>higher</p>
			</div>
							<div class="preporuka-data">
					<p>fuel tank</p>
					<p>
						<strong>
							<div class="data-good">
							3%							</div>
						</strong>
					</p>
					<p>larger</p>
				</div>
						<div class="preporuka-data">
				<p>boot</p>
				<p>
					<strong>
						<div class="data-good">
						8%						</div>
					</strong>
				</p>
				<p>larger</p>
			</div>
						<div class="preporuka-data">
				<p>boot ext.</p>
				<p>
					<strong>
						<div class="data-good">
						7%						</div>
					</strong>
				</p>
				<p>larger</p>
			</div>		
						<div class="preporuka-data">
				<p>price</p>
				<p>
					<strong>
						<div class="data-good">
						10%						</div>
					</strong>
				</p>
				<p>lower</p>
			</div>
		</div>
	</div>
</div><div class="box-wrap">
	<div class="preporuka-auto">
		<div class="preporuka-left">
			<div class="preporuka-slika">
				<a href="/compare/12927/12017/mazda-cx-5-skyactiv-d-2.2-175-4wd-VS-mitsubishi-outlander-2.2-di-d-4wd" title="compare selected cars">
					<img src="/resources/images/variant/1801/thumb_outlander_5.jpg" class="thumb-image">
				</a>
				<div class="preporuka-uporedi">
					<a href="/compare/12927/12017/mazda-cx-5-skyactiv-d-2.2-175-4wd-VS-mitsubishi-outlander-2.2-di-d-4wd" title="compare selected cars"><i class="fa fa-exchange" aria-hidden="true"></i></a>
				</div>
			</div>
		</div>
		<div class="preporuka-right">
			<div class="preporuka-naslov">
				<a href="/mitsubishi/2016/12017/mitsubishi-outlander-2.2-di-d-4wd" class="orange">Mitsubishi Outlander  2.2 DI-D 4WD</a>
			</div>
			<div class="preporuka-proizvodnja">
				produced from 2016. to 2018. &nbsp;			</div>
			<div class="preporuceni-stars">
				<div class="stars5"></div>
			</div>
			<div class="preporuceni-euro-ncap">
				EuroNCAP:
				88% score			</div>
		</div>
		<div class="preporuka-data-wrap">
			<div class="preporuka-data">
				<p>acceleration</p>
				<p>
					<strong>
						<div class="data-bad">
						14%						</div>
					</strong>
				</p>
				<p>slower</p>
			</div>
			<div class="preporuka-data">
				<p>consumption</p>
				<p>
					<strong>
						<div class="data-bad">
						2%						</div>
					</strong>
				</p>
				<p>higher</p>
			</div>
			<div class="preporuka-data">
				<p>power</p>
				<p>
					<strong>
						<div class="data-bad">
						17%						</div>
					</strong>
				</p>
				<p>lower</p>
			</div>
			<div class="preporuka-data">
				<p>length</p>
				<p>
					<strong>
						<div class="data-good">
						3%						</div>
					</strong>
				</p>
				<p>longer</p>
			</div>
			<div class="preporuka-data">
				<p>weight</p>
				<p>
					<strong>
						<div class="data-bad">
						5%						</div>
					</strong>
				</p>
				<p>higher</p>
			</div>
							<div class="preporuka-data">
					<p>fuel tank</p>
					<p>
						<strong>
							<div class="data-good">
							3%							</div>
						</strong>
					</p>
					<p>larger</p>
				</div>
						<div class="preporuka-data">
				<p>boot</p>
				<p>
					<strong>
						<div class="data-good">
						15%						</div>
					</strong>
				</p>
				<p>larger</p>
			</div>
						<div class="preporuka-data">
				<p>boot ext.</p>
				<p>
					<strong>
						<div class="data-good">
						59%						</div>
					</strong>
				</p>
				<p>smaller</p>
			</div>		
						<div class="preporuka-data">
				<p>price</p>
				<p>
					<strong>
						<div class="data-good">
						81%						</div>
					</strong>
				</p>
				<p>lower</p>
			</div>
		</div>
	</div>
</div><div class="box-wrap">
	<div class="preporuka-auto">
		<div class="preporuka-left">
			<div class="preporuka-slika">
				<a href="/compare/12927/12428/mazda-cx-5-skyactiv-d-2.2-175-4wd-VS-subaru-forester-2.0d" title="compare selected cars">
					<img src="/resources/images/variant/1860/thumb_forester_6.jpg" class="thumb-image">
				</a>
				<div class="preporuka-uporedi">
					<a href="/compare/12927/12428/mazda-cx-5-skyactiv-d-2.2-175-4wd-VS-subaru-forester-2.0d" title="compare selected cars"><i class="fa fa-exchange" aria-hidden="true"></i></a>
				</div>
			</div>
		</div>
		<div class="preporuka-right">
			<div class="preporuka-naslov">
				<a href="/subaru/2015/12428/subaru-forester-2.0d" class="orange">Subaru Forester  2.0D</a>
			</div>
			<div class="preporuka-proizvodnja">
				produced from 2015. to -			</div>
			<div class="preporuceni-stars">
				<div class="stars5"></div>
			</div>
			<div class="preporuceni-euro-ncap">
				EuroNCAP:
				91% score			</div>
		</div>
		<div class="preporuka-data-wrap">
			<div class="preporuka-data">
				<p>acceleration</p>
				<p>
					<strong>
						<div class="data-bad">
						14%						</div>
					</strong>
				</p>
				<p>slower</p>
			</div>
			<div class="preporuka-data">
				<p>consumption</p>
				<p>
					<strong>
						<div class="data-bad">
						9%						</div>
					</strong>
				</p>
				<p>higher</p>
			</div>
			<div class="preporuka-data">
				<p>power</p>
				<p>
					<strong>
						<div class="data-bad">
						17%						</div>
					</strong>
				</p>
				<p>lower</p>
			</div>
			<div class="preporuka-data">
				<p>length</p>
				<p>
					<strong>
						<div class="data-good">
						1%						</div>
					</strong>
				</p>
				<p>longer</p>
			</div>
			<div class="preporuka-data">
				<p>weight</p>
				<p>
					<strong>
						<div class="data-bad">
						4%						</div>
					</strong>
				</p>
				<p>higher</p>
			</div>
							<div class="preporuka-data">
					<p>fuel tank</p>
					<p>
						<strong>
							<div class="data-good">
							3%							</div>
						</strong>
					</p>
					<p>larger</p>
				</div>
						<div class="preporuka-data">
				<p>boot</p>
				<p>
					<strong>
						<div class="data-good">
						0%						</div>
					</strong>
				</p>
				<p>larger</p>
			</div>
						<div class="preporuka-data">
				<p>boot ext.</p>
				<p>
					<strong>
						<div class="data-good">
						2%						</div>
					</strong>
				</p>
				<p>smaller</p>
			</div>		
						<div class="preporuka-data">
				<p>price</p>
				<p>
					<strong>
						<div class="data-good">
						9%						</div>
					</strong>
				</p>
				<p>lower</p>
			</div>
		</div>
	</div>
</div><div class="box-wrap">
	<div class="preporuka-auto">
		<div class="preporuka-left">
			<div class="preporuka-slika">
				<a href="/compare/12927/13197/mazda-cx-5-skyactiv-d-2.2-175-4wd-VS-renault-koleos-dci-175-4wd" title="compare selected cars">
					<img src="/resources/images/variant/1971/thumb_koleos_3.jpg" class="thumb-image">
				</a>
				<div class="preporuka-uporedi">
					<a href="/compare/12927/13197/mazda-cx-5-skyactiv-d-2.2-175-4wd-VS-renault-koleos-dci-175-4wd" title="compare selected cars"><i class="fa fa-exchange" aria-hidden="true"></i></a>
				</div>
			</div>
		</div>
		<div class="preporuka-right">
			<div class="preporuka-naslov">
				<a href="/renault/2017/13197/renault-koleos-dci-175-4wd" class="orange">Renault Koleos  dCi 175 4WD</a>
			</div>
			<div class="preporuka-proizvodnja">
				produced from 2017. to 2020. &nbsp;			</div>
			<div class="preporuceni-stars">
				<div class="stars5"></div>
			</div>
			<div class="preporuceni-euro-ncap">
				EuroNCAP:
				90% score			</div>
		</div>
		<div class="preporuka-data-wrap">
			<div class="preporuka-data">
				<p>acceleration</p>
				<p>
					<strong>
						<div class="data-bad">
						18%						</div>
					</strong>
				</p>
				<p>slower</p>
			</div>
			<div class="preporuka-data">
				<p>consumption</p>
				<p>
					<strong>
						<div class="data-bad">
						2%						</div>
					</strong>
				</p>
				<p>higher</p>
			</div>
			<div class="preporuka-data">
				<p>power</p>
				<p>
					<strong>
						<div class="data-good">
						2%						</div>
					</strong>
				</p>
				<p>higher</p>
			</div>
			<div class="preporuka-data">
				<p>length</p>
				<p>
					<strong>
						<div class="data-good">
						3%						</div>
					</strong>
				</p>
				<p>longer</p>
			</div>
			<div class="preporuka-data">
				<p>weight</p>
				<p>
					<strong>
						<div class="data-bad">
						15%						</div>
					</strong>
				</p>
				<p>higher</p>
			</div>
							<div class="preporuka-data">
					<p>fuel tank</p>
					<p>
						<strong>
							<div class="data-good">
							3%							</div>
						</strong>
					</p>
					<p>larger</p>
				</div>
						<div class="preporuka-data">
				<p>boot</p>
				<p>
					<strong>
						<div class="data-good">
						5%						</div>
					</strong>
				</p>
				<p>larger</p>
			</div>
						<div class="preporuka-data">
				<p>boot ext.</p>
				<p>
					<strong>
						<div class="data-good">
						3%						</div>
					</strong>
				</p>
				<p>larger</p>
			</div>		
						<div class="preporuka-data">
				<p>price</p>
				<p>
					<strong>
						<div class="data-good">
						20%						</div>
					</strong>
				</p>
				<p>lower</p>
			</div>
		</div>
	</div>
</div><div>
	<h2><i class="fa fa-search" aria-hidden="true"></i>Check a car by its VIN number</h2>
</div>
<div class="box-wrap" style="height: 250px;">
	<div
	  data-cvaff
	  data-locale="en-EU"
	  data-a="automanijak"
	  data-b="81ec5429"
	  style="width: 100%; height: 100%">
	</div>
</div><div>
	<h2><i class="fa fa-user" aria-hidden="true"></i>Visitor comments</h2>
</div>
<div class="box-wrap">
	<div class="fb-comments" data-href="https://www.automaniac.org/mazda/2017/12927/mazda-cx-5-skyactiv-d-2.2-175-4wd" data-width="670" data-numposts="5"></div>
</div>
	</div>
			</div>
		</div>
		<div id="footer">
			<div class="footer-wrap">
				Copyright &copy; 2015 - 2022 automaniac.org - All rights reserved.
				</br>
				Powered by <strong>Gama1 Solutions</strong>
				<div class="footer-links">
					<div class="footer-menu">
						<ul>
							<li><a href="/">HOME page</a></li>
							<li><a href="/about">about</a></li>
							<li><a href="/terms">terms & conditions</a></li>
							<li><a href="/legal">legal note</a></li>
							<li><a href="/faq">frequent questions (FAQ)</a></li>
							<li><a href="/contact">contact us</a></li>
						</ul>
					</div>

					<div class='footer-auto-brend'><ul><li><a href='/make/alfa-romeo'>Alfa Romeo</a></li><li><a href='/make/audi'>Audi</a></li><li><a href='/make/bmw'>BMW</a></li><li><a href='/make/chery'>Chery</a></li><li><a href='/make/chevrolet'>Chevrolet</a></li><li><a href='/make/chrysler'>Chrysler</a></li></ul></div><div class='footer-auto-brend'><ul><li><a href='/make/citroen'>Citroen</a></li><li><a href='/make/dacia'>Dacia</a></li><li><a href='/make/daihatsu'>Daihatsu</a></li><li><a href='/make/dodge'>Dodge</a></li><li><a href='/make/fiat'>FIAT</a></li><li><a href='/make/ford'>Ford</a></li></ul></div><div class='footer-auto-brend'><ul><li><a href='/make/honda'>Honda</a></li><li><a href='/make/hyundai'>Hyundai</a></li><li><a href='/make/infiniti'>Infiniti</a></li><li><a href='/make/jaguar'>Jaguar</a></li><li><a href='/make/jeep'>Jeep</a></li><li><a href='/make/kia'>KIA</a></li></ul></div><div class='footer-auto-brend'><ul><li><a href='/make/lada'>Lada</a></li><li><a href='/make/lancia'>Lancia</a></li><li><a href='/make/land-rover'>Land Rover</a></li><li><a href='/make/lexus'>Lexus</a></li><li><a href='/make/mazda'>Mazda</a></li><li><a href='/make/mercedes-benz'>Mercedes</a></li></ul></div><div class='footer-auto-brend'><ul><li><a href='/make/mini'>Mini</a></li><li><a href='/make/mitsubishi'>Mitsubishi</a></li><li><a href='/make/nissan'>Nissan</a></li><li><a href='/make/opel'>Opel</a></li><li><a href='/make/peugeot'>Peugeot</a></li><li><a href='/make/porsche'>Porsche</a></li></ul></div><div class='footer-auto-brend'><ul><li><a href='/make/renault'>Renault</a></li><li><a href='/make/rover'>Rover</a></li><li><a href='/make/saab'>SAAB</a></li><li><a href='/make/ssangyong'>SSangYong</a></li><li><a href='/make/seat'>Seat</a></li><li><a href='/make/smart'>Smart</a></li></ul></div><div class='footer-auto-brend'><ul><li><a href='/make/subaru'>Subaru</a></li><li><a href='/make/suzuki'>Suzuki</a></li><li><a href='/make/tesla'>Tesla</a></li><li><a href='/make/toyota'>Toyota</a></li><li><a href='/make/volkswagen'>Volkswagen</a></li><li><a href='/make/volvo'>Volvo</a></li></ul></div><div class='footer-auto-brend'><ul><li><a href='/make/zastava'>Zastava</a></li><li><a href='/make/škoda'>Škoda</a></li></ul></div>
				</div>
			</div>
		</div>
	</body>
</html>`
