
import java.util.HashMap;
import java.io.File;


public class ConComp {
	public static void main(String[] args) {
		System.out.println("Number of command line parameters: " + args.length);
		
		java.util.HashMap<String, Long> lengths = new java.util.HashMap<>();
		for(String fname: args) {
			File file = new File(fname);
			lengths.put(fname, file.length());
		}
		
		String longest="";
		long maxlen=0;
		for(String fname: lengths.keySet()) {
			if(lengths.get(fname) > maxlen) {
				maxlen = lengths.get(fname);
				longest = fname;
			}
		}
		
		System.out.println("File " + longest + " is the biggest at " + maxlen + " bytes");
		
	}
}