#import "NSString+JSONSorted.h"
#import "NSArray+JSONSorted.h"
#import "NSDictionary+JSONSorted.h"
@implementation NSString (JSONSorted) 
- (NSString *)sortedJSONString {
       if ([self hasPrefix:@"{"]) {
           //转成NSDictionary
       } else if ([self hasPrefix:@"["]) {
           //转成NSArray
       } else {
           return self;
       }
}
@end